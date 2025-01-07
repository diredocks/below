import { createSignal, For } from "solid-js";


const CommentList = ({ comments }) => {


  return (
    <>
      <div class="text-lg font-bold">评论（{comments().length}）</div>
      <ul class="mt-4">
        <For each={comments()}>
          {(comment) => (
            <li class="border-b py-4 first:pt-0 last:border-none">
              <div class="flex flex-row items-center gap-2">
                <p class="text-base font-semibold">{comment.name}</p>
                <p class="text-sm text-slate-500">{comment.createdAt}</p>
              </div>
              <p class="text-base break-all mt-1">{comment.content}</p>
            </li>
          )}
        </For>
      </ul>
    </>
  );
};

export default CommentList;
