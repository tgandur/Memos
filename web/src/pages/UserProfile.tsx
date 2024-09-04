import { Button } from "@mui/joy";
import copy from "copy-to-clipboard";
import dayjs from "dayjs";
import { ArrowDownIcon, ExternalLinkIcon } from "lucide-react";
import { useEffect, useState } from "react";
import { toast } from "react-hot-toast";
import { useParams } from "react-router-dom";
import Empty from "@/components/Empty";
import MemoFilters from "@/components/MemoFilters";
import MemoView from "@/components/MemoView";
import MobileHeader from "@/components/MobileHeader";
import UserAvatar from "@/components/UserAvatar";
import { DEFAULT_LIST_MEMOS_PAGE_SIZE } from "@/helpers/consts";
import useLoading from "@/hooks/useLoading";
import { useMemoFilterStore, useMemoList, useMemoStore, useUserStore } from "@/store/v1";
import { User } from "@/types/proto/api/v1/user_service";
import { useTranslate } from "@/utils/i18n";

const UserProfile = () => {
  const t = useTranslate();
  const params = useParams();
  const userStore = useUserStore();
  const loadingState = useLoading();
  const [user, setUser] = useState<User>();
  const memoStore = useMemoStore();
  const memoList = useMemoList();
  const memoFilterStore = useMemoFilterStore();
  const [isRequesting, setIsRequesting] = useState(true);
  const [nextPageToken, setNextPageToken] = useState<string>("");
  const sortedMemos = memoList.value
    .sort((a, b) => dayjs(b.displayTime).unix() - dayjs(a.displayTime).unix())
    .sort((a, b) => Number(b.pinned) - Number(a.pinned));

  useEffect(() => {
    const username = params.username;
    if (!username) {
      throw new Error("username is required");
    }

    userStore
      .searchUsers(`username == "${username}"`)
      .then((users) => {
        if (users.length !== 1) {
          throw new Error("User not found");
        }
        const user = users[0];
        setUser(user);
        loadingState.setFinish();
      })
      .catch((error) => {
        console.error(error);
        toast.error(t("message.user-not-found"));
      });
  }, [params.username]);

  useEffect(() => {
    if (!user) {
      return;
    }

    memoList.reset();
    fetchMemos("");
  }, [user, memoFilterStore.filters]);

  const fetchMemos = async (nextPageToken: string) => {
    if (!user) {
      return;
    }

    setIsRequesting(true);
    const filters = [`creator == "${user.name}"`, `row_status == "NORMAL"`, `order_by_pinned == true`];
    const contentSearch: string[] = [];
    const tagSearch: string[] = [];
    for (const filter of memoFilterStore.filters) {
      if (filter.factor === "contentSearch") {
        contentSearch.push(`"${filter.value}"`);
      } else if (filter.factor === "tagSearch") {
        tagSearch.push(`"${filter.value}"`);
      }
    }
    if (contentSearch.length > 0) {
      filters.push(`content_search == [${contentSearch.join(", ")}]`);
    }
    if (tagSearch.length > 0) {
      filters.push(`tag_search == [${tagSearch.join(", ")}]`);
    }
    const response = await memoStore.fetchMemos({
      pageSize: DEFAULT_LIST_MEMOS_PAGE_SIZE,
      filter: filters.join(" && "),
      pageToken: nextPageToken,
    });
    setIsRequesting(false);
    setNextPageToken(response.nextPageToken);
  };

  const handleCopyProfileLink = () => {
    if (!user) {
      return;
    }

    copy(`${window.location.origin}` + (window as any).globalConfig.BaseUrl + `/u/${encodeURIComponent(user.username)}`);
    toast.success(t("message.copied"));
  };

  return (
    <section className="w-full max-w-5xl min-h-full flex flex-col justify-start items-center sm:pt-3 md:pt-6 pb-8">
      <MobileHeader />
      <div className="w-full px-4 sm:px-6 flex flex-col justify-start items-center">
        {!loadingState.isLoading &&
          (user ? (
            <>
              <div className="my-4 w-full flex justify-end items-center gap-2">
                <Button
                  color="neutral"
                  variant="outlined"
                  endDecorator={<ExternalLinkIcon className="w-4 h-auto opacity-60" />}
                  onClick={handleCopyProfileLink}
                >
                  {t("common.share")}
                </Button>
              </div>
              <div className="w-full flex flex-col justify-start items-start pt-4 pb-8 px-3">
                <UserAvatar className="!w-16 !h-16 drop-shadow rounded-3xl" avatarUrl={user?.avatarUrl} />
                <div className="mt-2 w-auto max-w-[calc(100%-6rem)] flex flex-col justify-center items-start">
                  <p className="w-full text-3xl text-black leading-tight opacity-80 dark:text-gray-200 truncate">
                    {user.nickname || user.username}
                  </p>
                  <p className="w-full text-gray-500 leading-snug opacity-80 dark:text-gray-400 whitespace-pre-wrap truncate line-clamp-6">
                    {user.description}
                  </p>
                </div>
              </div>
              <MemoFilters />
              {sortedMemos.map((memo) => (
                <MemoView key={`${memo.name}-${memo.displayTime}`} memo={memo} showVisibility showPinned compact />
              ))}
              {nextPageToken && (
                <div className="w-full flex flex-row justify-center items-center my-4">
                  <Button
                    variant="plain"
                    color="neutral"
                    loading={isRequesting}
                    endDecorator={<ArrowDownIcon className="w-4 h-auto" />}
                    onClick={() => fetchMemos(nextPageToken)}
                  >
                    {t("memo.load-more")}
                  </Button>
                </div>
              )}
              {!nextPageToken && sortedMemos.length === 0 && (
                <div className="w-full mt-12 mb-8 flex flex-col justify-center items-center italic">
                  <Empty />
                  <p className="mt-2 text-gray-600 dark:text-gray-400">{t("message.no-data")}</p>
                </div>
              )}
            </>
          ) : (
            <p>Not found</p>
          ))}
      </div>
    </section>
  );
};

export default UserProfile;
