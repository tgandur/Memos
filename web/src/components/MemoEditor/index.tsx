import { Select, Option, Button, Divider } from "@mui/joy";
import { isEqual } from "lodash-es";
import { SendIcon } from "lucide-react";
import React, { useEffect, useMemo, useRef, useState } from "react";
import { toast } from "react-hot-toast";
import { useTranslation } from "react-i18next";
import useLocalStorage from "react-use/lib/useLocalStorage";
import { memoServiceClient } from "@/grpcweb";
import { TAB_SPACE_WIDTH } from "@/helpers/consts";
import { isValidUrl } from "@/helpers/utils";
import useAsyncEffect from "@/hooks/useAsyncEffect";
import useCurrentUser from "@/hooks/useCurrentUser";
import { useMemoStore, useResourceStore, useUserStore, useWorkspaceSettingStore } from "@/store/v1";
import { MemoRelation, MemoRelation_Type } from "@/types/proto/api/v1/memo_relation_service";
import { Memo, Visibility } from "@/types/proto/api/v1/memo_service";
import { Resource } from "@/types/proto/api/v1/resource_service";
import { UserSetting } from "@/types/proto/api/v1/user_service";
import { WorkspaceMemoRelatedSetting } from "@/types/proto/api/v1/workspace_setting_service";
import { WorkspaceSettingKey } from "@/types/proto/store/workspace_setting";
import { useTranslate } from "@/utils/i18n";
import { convertVisibilityFromString, convertVisibilityToString } from "@/utils/memo";
import VisibilityIcon from "../VisibilityIcon";
import AddMemoRelationPopover from "./ActionButton/AddMemoRelationPopover";
import MarkdownMenu from "./ActionButton/MarkdownMenu";
import TagSelector from "./ActionButton/TagSelector";
import UploadResourceButton from "./ActionButton/UploadResourceButton";
import Editor, { EditorRefActions } from "./Editor";
import RelationListView from "./RelationListView";
import ResourceListView from "./ResourceListView";
import { handleEditorKeydownWithMarkdownShortcuts, hyperlinkHighlightedText } from "./handlers";
import { MemoEditorContext } from "./types";

export interface Props {
  className?: string;
  cacheKey?: string;
  placeholder?: string;
  // The name of the memo to be edited.
  memoName?: string;
  // The name of the parent memo if the memo is a comment.
  parentMemoName?: string;
  autoFocus?: boolean;
  onConfirm?: (memoName: string) => void;
  onCancel?: () => void;
}

interface State {
  memoVisibility: Visibility;
  resourceList: Resource[];
  relationList: MemoRelation[];
  isUploadingResource: boolean;
  isRequesting: boolean;
  isComposing: boolean;
}

const MemoEditor = (props: Props) => {
  const { className, cacheKey, memoName, parentMemoName, autoFocus, onConfirm, onCancel } = props;
  const t = useTranslate();
  const { i18n } = useTranslation();
  const workspaceSettingStore = useWorkspaceSettingStore();
  const userStore = useUserStore();
  const memoStore = useMemoStore();
  const resourceStore = useResourceStore();
  const currentUser = useCurrentUser();
  const [state, setState] = useState<State>({
    memoVisibility: Visibility.PRIVATE,
    resourceList: [],
    relationList: [],
    isUploadingResource: false,
    isRequesting: false,
    isComposing: false,
  });
  const [displayTime, setDisplayTime] = useState<Date | undefined>();
  const [hasContent, setHasContent] = useState<boolean>(false);
  const editorRef = useRef<EditorRefActions>(null);
  const userSetting = userStore.userSetting as UserSetting;
  const contentCacheKey = `${currentUser.name}-${cacheKey || ""}`;
  const [contentCache, setContentCache] = useLocalStorage<string>(contentCacheKey, "");
  const referenceRelations = memoName
    ? state.relationList.filter(
        (relation) => relation.memo === memoName && relation.relatedMemo !== memoName && relation.type === MemoRelation_Type.REFERENCE,
      )
    : state.relationList.filter((relation) => relation.type === MemoRelation_Type.REFERENCE);
  const workspaceMemoRelatedSetting =
    workspaceSettingStore.getWorkspaceSettingByKey(WorkspaceSettingKey.MEMO_RELATED)?.memoRelatedSetting ||
    WorkspaceMemoRelatedSetting.fromPartial({});

  useEffect(() => {
    editorRef.current?.setContent(contentCache || "");
  }, []);

  useEffect(() => {
    if (autoFocus) {
      handleEditorFocus();
    }
  }, [autoFocus]);

  useEffect(() => {
    let visibility = userSetting.memoVisibility;
    if (workspaceMemoRelatedSetting.disallowPublicVisibility && visibility === "PUBLIC") {
      visibility = "PRIVATE";
    }
    setState((prevState) => ({
      ...prevState,
      memoVisibility: convertVisibilityFromString(visibility),
    }));
  }, [userSetting.memoVisibility, workspaceMemoRelatedSetting.disallowPublicVisibility]);

  useAsyncEffect(async () => {
    if (!memoName) {
      return;
    }

    const memo = await memoStore.getOrFetchMemoByName(memoName);
    if (memo) {
      handleEditorFocus();
      setState((prevState) => ({
        ...prevState,
        memoVisibility: memo.visibility,
        resourceList: memo.resources,
        relationList: memo.relations,
      }));
      setDisplayTime(memo.displayTime);
      if (!contentCache) {
        editorRef.current?.setContent(memo.content ?? "");
      }
    }
  }, [memoName]);

  const handleCompositionStart = () => {
    setState((prevState) => ({
      ...prevState,
      isComposing: true,
    }));
  };

  const handleCompositionEnd = () => {
    setState((prevState) => ({
      ...prevState,
      isComposing: false,
    }));
  };

  const handleKeyDown = (event: React.KeyboardEvent) => {
    if (!editorRef.current) {
      return;
    }

    const isMetaKey = event.ctrlKey || event.metaKey;
    if (isMetaKey) {
      if (event.key === "Enter") {
        void handleSaveBtnClick();
        return;
      }

      handleEditorKeydownWithMarkdownShortcuts(event, editorRef.current);
    }
    if (event.key === "Tab" && !state.isComposing) {
      event.preventDefault();
      const tabSpace = " ".repeat(TAB_SPACE_WIDTH);
      const cursorPosition = editorRef.current.getCursorPosition();
      const selectedContent = editorRef.current.getSelectedContent();
      editorRef.current.insertText(tabSpace);
      if (selectedContent) {
        editorRef.current.setCursorPosition(cursorPosition + TAB_SPACE_WIDTH);
      }
      return;
    }
  };

  const handleMemoVisibilityChange = (visibility: Visibility) => {
    setState((prevState) => ({
      ...prevState,
      memoVisibility: visibility,
    }));
  };

  const handleSetResourceList = (resourceList: Resource[]) => {
    setState((prevState) => ({
      ...prevState,
      resourceList,
    }));
  };

  const handleSetRelationList = (relationList: MemoRelation[]) => {
    setState((prevState) => ({
      ...prevState,
      relationList,
    }));
  };

  const handleUploadResource = async (file: File) => {
    setState((state) => {
      return {
        ...state,
        isUploadingResource: true,
      };
    });

    const { name: filename, size, type } = file;
    const buffer = new Uint8Array(await file.arrayBuffer());

    try {
      const resource = await resourceStore.createResource({
        resource: Resource.fromPartial({
          filename,
          size,
          type,
          content: buffer,
        }),
      });
      setState((state) => {
        return {
          ...state,
          isUploadingResource: false,
        };
      });
      return resource;
    } catch (error: any) {
      console.error(error);
      toast.error(error.details);
    }
  };

  const uploadMultiFiles = async (files: FileList) => {
    const uploadedResourceList: Resource[] = [];
    for (const file of files) {
      const resource = await handleUploadResource(file);
      if (resource) {
        uploadedResourceList.push(resource);
        if (memoName) {
          await resourceStore.updateResource({
            resource: Resource.fromPartial({
              name: resource.name,
              memo: memoName,
            }),
            updateMask: ["memo"],
          });
        }
      }
    }
    if (uploadedResourceList.length > 0) {
      setState((prevState) => ({
        ...prevState,
        resourceList: [...prevState.resourceList, ...uploadedResourceList],
      }));
    }
  };

  const handleDropEvent = async (event: React.DragEvent) => {
    if (event.dataTransfer && event.dataTransfer.files.length > 0) {
      event.preventDefault();
      await uploadMultiFiles(event.dataTransfer.files);
    }
  };

  const handlePasteEvent = async (event: React.ClipboardEvent) => {
    if (event.clipboardData && event.clipboardData.files.length > 0) {
      event.preventDefault();
      await uploadMultiFiles(event.clipboardData.files);
    } else if (
      editorRef.current != null &&
      editorRef.current.getSelectedContent().length != 0 &&
      isValidUrl(event.clipboardData.getData("Text"))
    ) {
      event.preventDefault();
      hyperlinkHighlightedText(editorRef.current, event.clipboardData.getData("Text"));
    }
  };

  const handleContentChange = (content: string) => {
    setHasContent(content !== "");
    if (content !== "") {
      setContentCache(content);
    } else {
      localStorage.removeItem(contentCacheKey);
    }
  };

  const handleSaveBtnClick = async () => {
    if (state.isRequesting) {
      return;
    }

    setState((state) => {
      return {
        ...state,
        isRequesting: true,
      };
    });
    const content = editorRef.current?.getContent() ?? "";
    try {
      // Update memo.
      if (memoName) {
        const prevMemo = await memoStore.getOrFetchMemoByName(memoName);
        if (prevMemo) {
          const updateMask = ["content", "visibility"];
          const memoPatch: Partial<Memo> = {
            name: prevMemo.name,
            content,
            visibility: state.memoVisibility,
          };
          if (!isEqual(displayTime, prevMemo.displayTime)) {
            updateMask.push("display_time");
            memoPatch.displayTime = displayTime;
          }
          const memo = await memoStore.updateMemo(memoPatch, updateMask);
          await memoServiceClient.setMemoResources({
            name: memo.name,
            resources: state.resourceList,
          });
          await memoServiceClient.setMemoRelations({
            name: memo.name,
            relations: state.relationList,
          });
          await memoStore.getOrFetchMemoByName(memo.name, { skipCache: true });
          if (onConfirm) {
            onConfirm(memo.name);
          }
        }
      } else {
        // Create memo or memo comment.
        const request = !parentMemoName
          ? memoStore.createMemo({
              content,
              visibility: state.memoVisibility,
            })
          : memoServiceClient
              .createMemoComment({
                name: parentMemoName,
                comment: {
                  content,
                  visibility: state.memoVisibility,
                },
              })
              .then((memo) => memo);
        const memo = await request;
        await memoServiceClient.setMemoResources({
          name: memo.name,
          resources: state.resourceList,
        });
        await memoServiceClient.setMemoRelations({
          name: memo.name,
          relations: state.relationList,
        });
        await memoStore.getOrFetchMemoByName(memo.name, { skipCache: true });
        if (onConfirm) {
          onConfirm(memo.name);
        }
      }
      editorRef.current?.setContent("");
    } catch (error: any) {
      console.error(error);
      toast.error(error.details);
    }

    localStorage.removeItem(contentCacheKey);
    setState((state) => {
      return {
        ...state,
        isRequesting: false,
        resourceList: [],
        relationList: [],
      };
    });
  };

  const handleCancelBtnClick = () => {
    localStorage.removeItem(contentCacheKey);

    if (onCancel) {
      onCancel();
    }
  };

  const handleEditorFocus = () => {
    editorRef.current?.focus();
  };

  const editorConfig = useMemo(
    () => ({
      className: "",
      initialContent: "",
      placeholder: props.placeholder ?? t("editor.any-thoughts"),
      onContentChange: handleContentChange,
      onPaste: handlePasteEvent,
    }),
    [i18n.language],
  );

  const allowSave = (hasContent || state.resourceList.length > 0) && !state.isUploadingResource && !state.isRequesting;

  /**
   * Format a date to a string yyyy-mm-ddThh:mm:ss
   * @param date The date to convert
   * @returns A string representation of the date
   */
  const toCalendarFormatString = (date: Date): string => {
    return new Date(date.getTime() - date.getTimezoneOffset() * 60000).toISOString().split(".")[0];
  };

  return (
    <MemoEditorContext.Provider
      value={{
        resourceList: state.resourceList,
        relationList: state.relationList,
        setResourceList: (resourceList: Resource[]) => {
          setState((prevState) => ({
            ...prevState,
            resourceList,
          }));
        },
        setRelationList: (relationList: MemoRelation[]) => {
          setState((prevState) => ({
            ...prevState,
            relationList,
          }));
        },
        memoName,
      }}
    >
      <div
        className={`${
          className ?? ""
        } relative w-full flex flex-col justify-start items-start bg-white dark:bg-zinc-800 px-4 pt-4 rounded-lg border border-gray-200 dark:border-zinc-700`}
        tabIndex={0}
        onKeyDown={handleKeyDown}
        onDrop={handleDropEvent}
        onFocus={handleEditorFocus}
        onCompositionStart={handleCompositionStart}
        onCompositionEnd={handleCompositionEnd}
      >
        {memoName && displayTime && (
          <div className="relative text-sm">
            <span className="cursor-pointer text-gray-400 dark:text-gray-500">{displayTime.toLocaleString()}</span>
            <input
              className="inset-0 absolute z-1 opacity-0"
              type="datetime-local"
              value={toCalendarFormatString(displayTime)}
              onFocus={(e: any) => e.target.showPicker()}
              onChange={(e) => setDisplayTime(new Date(e.target.value))}
            />
          </div>
        )}
        <Editor ref={editorRef} {...editorConfig} />
        <ResourceListView resourceList={state.resourceList} setResourceList={handleSetResourceList} />
        <RelationListView relationList={referenceRelations} setRelationList={handleSetRelationList} />
        <div className="relative w-full flex flex-row justify-between items-center pt-2" onFocus={(e) => e.stopPropagation()}>
          <div className="flex flex-row justify-start items-center opacity-80 dark:opacity-60">
            <TagSelector editorRef={editorRef} />
            <MarkdownMenu editorRef={editorRef} />
            <UploadResourceButton />
            <AddMemoRelationPopover editorRef={editorRef} />
          </div>
        </div>
        <Divider className="!mt-2" />
        <div className="w-full flex flex-row justify-between items-center py-3 dark:border-t-zinc-500">
          <div className="relative flex flex-row justify-start items-center" onFocus={(e) => e.stopPropagation()}>
            <Select
              variant="plain"
              value={state.memoVisibility}
              startDecorator={<VisibilityIcon visibility={state.memoVisibility} />}
              onChange={(_, visibility) => {
                if (visibility) {
                  handleMemoVisibilityChange(visibility);
                }
              }}
            >
              {[Visibility.PRIVATE, Visibility.PROTECTED, Visibility.PUBLIC].map((item) => (
                <Option key={item} value={item} className="whitespace-nowrap">
                  {t(`memo.visibility.${convertVisibilityToString(item).toLowerCase()}` as any)}
                </Option>
              ))}
            </Select>
          </div>
          <div className="shrink-0 flex flex-row justify-end items-center gap-2">
            {props.onCancel && (
              <Button className="!font-normal" color="neutral" variant="plain" loading={state.isRequesting} onClick={handleCancelBtnClick}>
                {t("common.cancel")}
              </Button>
            )}
            <Button
              className="!font-normal"
              disabled={!allowSave}
              loading={state.isRequesting}
              endDecorator={<SendIcon className="w-4 h-auto" />}
              onClick={handleSaveBtnClick}
            >
              {t("editor.save")}
            </Button>
          </div>
        </div>
      </div>
    </MemoEditorContext.Provider>
  );
};

export default MemoEditor;
