import dayjs from "dayjs";
import { useMemo } from "react";
import { HomeSidebar, HomeSidebarDrawer } from "@/components/HomeSidebar";
import MemoEditor from "@/components/MemoEditor";
import MemoView from "@/components/MemoView";
import MobileHeader from "@/components/MobileHeader";
import PagedMemoList from "@/components/PagedMemoList";
import useCurrentUser from "@/hooks/useCurrentUser";
import useResponsiveWidth from "@/hooks/useResponsiveWidth";
import { useMemoFilterStore, useUserStore } from "@/store/v1";
import { Direction, State } from "@/types/proto/api/v1/common";
import { Memo } from "@/types/proto/api/v1/memo_service";
import { cn } from "@/utils";

const Home = () => {
  const { md, lg } = useResponsiveWidth();
  const user = useCurrentUser();
  const userStore = useUserStore();
  const memoFilterStore = useMemoFilterStore();
  const selectedShortcut = userStore.shortcuts.find((shortcut) => shortcut.id === memoFilterStore.shortcut);

  const memoListFilter = useMemo(() => {
    const conditions = [];
    const contentSearch: string[] = [];
    const tagSearch: string[] = [];
    for (const filter of memoFilterStore.filters) {
      if (filter.factor === "contentSearch") {
        contentSearch.push(`"${filter.value}"`);
      } else if (filter.factor === "tagSearch") {
        tagSearch.push(`"${filter.value}"`);
      } else if (filter.factor === "property.hasLink") {
        conditions.push(`has_link == true`);
      } else if (filter.factor === "property.hasTaskList") {
        conditions.push(`has_task_list == true`);
      } else if (filter.factor === "property.hasCode") {
        conditions.push(`has_code == true`);
      } else if (filter.factor === "displayTime") {
        const filterDate = new Date(filter.value);
        const filterUtcTimestamp = filterDate.getTime() + filterDate.getTimezoneOffset() * 60 * 1000;
        const timestampAfter = filterUtcTimestamp / 1000;
        conditions.push(`display_time_after == ${timestampAfter}`);
        conditions.push(`display_time_before == ${timestampAfter + 60 * 60 * 24}`);
      }
    }
    if (contentSearch.length > 0) {
      conditions.push(`content_search == [${contentSearch.join(", ")}]`);
    }
    if (tagSearch.length > 0) {
      conditions.push(`tag_search == [${tagSearch.join(", ")}]`);
    }
    return conditions.join(" && ");
  }, [user, memoFilterStore.filters, memoFilterStore.orderByTimeAsc]);

  return (
    <section className="@container w-full min-h-full flex flex-col justify-start items-center">
      {!md && (
        <MobileHeader>
          <HomeSidebarDrawer />
        </MobileHeader>
      )}
      <div className={cn("w-full min-h-full flex flex-row justify-start items-start")}>
        {md && (
          <div
            className={cn(
              "sticky top-0 left-0 shrink-0 h-[100svh] transition-all",
              "border-r border-gray-200 dark:border-zinc-800",
              lg ? "px-5 w-72" : "px-4 w-56",
            )}
          >
            <HomeSidebar className={cn("py-6")} />
          </div>
        )}
        <div className={cn("w-full mx-auto px-4 sm:px-6 sm:pt-3 md:pt-6 pb-8", md && "max-w-3xl")}>
          <MemoEditor className="mb-2" cacheKey="home-memo-editor" />
          <div className="flex flex-col justify-start items-start w-full max-w-full">
            <PagedMemoList
              renderer={(memo: Memo) => <MemoView key={`${memo.name}-${memo.displayTime}`} memo={memo} showVisibility showPinned compact />}
              listSort={(memos: Memo[]) =>
                memos
                  .filter((memo) => memo.state === State.NORMAL)
                  .sort((a, b) =>
                    memoFilterStore.orderByTimeAsc
                      ? dayjs(a.displayTime).unix() - dayjs(b.displayTime).unix()
                      : dayjs(b.displayTime).unix() - dayjs(a.displayTime).unix(),
                  )
                  .sort((a, b) => Number(b.pinned) - Number(a.pinned))
              }
              owner={user.name}
              direction={memoFilterStore.orderByTimeAsc ? Direction.ASC : Direction.DESC}
              filter={selectedShortcut?.filter || ""}
              oldFilter={memoListFilter}
            />
          </div>
        </div>
      </div>
    </section>
  );
};

export default Home;
