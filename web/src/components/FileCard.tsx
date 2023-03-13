import dayjs from "dayjs";
import { useState } from "react";
import { useTranslation } from "react-i18next";
import { toast } from "react-hot-toast";
import { useResourceStore } from "../store/module";
import Icon from "../components/Icon";
import copy from "copy-to-clipboard";
import { getResourceUrl } from "../utils/resource";
import showPreviewImageDialog from "../components/PreviewImageDialog";
import Dropdown from "./base/Dropdown";
import FileCover from "./FileCover";
import { showCommonDialog } from "../components/Dialog/CommonDialog";
import showChangeResourceFilenameDialog from "../components/ChangeResourceFilenameDialog";

import "../less/file-card.less";

interface FileProps {
  resource: Resource;
  handlecheck: any;
  handleUncheck: any;
}

const FileCard = ({ resource, handlecheck, handleUncheck }: FileProps) => {
  const [isSelected, setIsSelected] = useState<boolean>(false);
  const resourceStore = useResourceStore();
  const resources = resourceStore.state.resources;
  const { t } = useTranslation();

  const handleRenameBtnClick = (resource: Resource) => {
    showChangeResourceFilenameDialog(resource.id, resource.filename);
  };

  const handleDeleteResourceBtnClick = (resource: Resource) => {
    let warningText = t("resources.warning-text");
    if (resource.linkedMemoAmount > 0) {
      warningText = warningText + `\n${t("resources.linked-amount")}: ${resource.linkedMemoAmount}`;
    }

    showCommonDialog({
      title: t("resources.delete-resource"),
      content: warningText,
      style: "warning",
      dialogName: "delete-resource-dialog",
      onConfirm: async () => {
        await resourceStore.deleteResourceById(resource.id);
      },
    });
  };

  const handlePreviewBtnClick = (resource: Resource) => {
    const resourceUrl = getResourceUrl(resource);
    if (resource.type.startsWith("image")) {
      showPreviewImageDialog(
        resources.filter((r) => r.type.startsWith("image")).map((r) => getResourceUrl(r)),
        resources.findIndex((r) => r.id === resource.id)
      );
    } else {
      window.open(resourceUrl);
    }
  };

  const handleCopyResourceLinkBtnClick = (resource: Resource) => {
    const url = getResourceUrl(resource);
    copy(url);
    toast.success(t("message.succeed-copy-resource-link"));
  };

  const handleSelectBtnClick = () => {
    if (isSelected) {
      handlecheck();
    } else {
      handleUncheck();
    }
    setIsSelected(!isSelected);
  };

  return (
    <div className="resource-card">
      <div className="btns-container">
        <div onClick={() => handleSelectBtnClick()}>
          {isSelected ? (
            <Icon.CheckCircle2 className="m-2 text-gray-500 hover:text-black" />
          ) : (
            <Icon.Circle className="resource-checkbox" />
          )}
        </div>

        <Dropdown
          className="more-action-btn"
          actionsClassName="!w-28"
          actions={
            <>
              <button
                className="w-full text-left text-sm leading-6 py-1 px-3 cursor-pointer rounded hover:bg-gray-100 dark:hover:bg-zinc-600"
                onClick={() => handlePreviewBtnClick(resource)}
              >
                {t("resources.preview")}
              </button>
              <button
                className="w-full text-left text-sm leading-6 py-1 px-3 cursor-pointer rounded hover:bg-gray-100 dark:hover:bg-zinc-600"
                onClick={() => handleCopyResourceLinkBtnClick(resource)}
              >
                {t("resources.copy-link")}
              </button>
              <button
                className="w-full text-left text-sm leading-6 py-1 px-3 cursor-pointer rounded hover:bg-gray-100 dark:hover:bg-zinc-600"
                onClick={() => handleRenameBtnClick(resource)}
              >
                {t("resources.rename")}
              </button>
              <button
                className="w-full text-left text-sm leading-6 py-1 px-3 cursor-pointer rounded text-red-600 hover:bg-gray-100 dark:hover:bg-zinc-600"
                onClick={() => handleDeleteResourceBtnClick(resource)}
              >
                {t("common.delete")}
              </button>
            </>
          }
        />
      </div>
      <FileCover resource={resource} />
      <div>
        <div className="text-base overflow-ellipsis text-center">{resource.filename}</div>
        <div className="text-sm text-gray-400 text-center">{dayjs(resource.createdTs).locale("en").format("YYYY/MM/DD HH:mm:ss")}</div>
      </div>
    </div>
  );
};

export default FileCard;
