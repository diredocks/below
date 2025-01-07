import { For } from "solid-js";
import { formatDateTimeTo } from "./DateAndTime";

const CommentList = ({ comments }) => {
  return (
    <>
      <p class="comment-list-header text-lg font-bold">评论（{comments().length}）</p>
      <ul class="mt-4">
        <For each={comments()}>
          {(comment) => (
            <li class="comment-item border-b py-4 first:pt-0 last:border-none">
              <div class="comment-header flex flex-row items-center gap-2">
                <p class="comment-name text-base font-semibold">{comment.name}</p>
                <p class="comment-time text-sm text-slate-500">
                  {formatDateTimeTo(comment.CreatedAt)}
                </p>
              </div>
              <p class="comment-content text-base break-all mt-1">{comment.content}</p>
            </li>
          )}
        </For>
      </ul>
    </>
  );
};

export default CommentList;
