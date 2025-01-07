import { createSignal, For } from "solid-js";


const CommentList = () => {

  const [comments] = createSignal([
    { name: "Kytoki", content: "Hello there!", createdAt: "2025/1/4 12:34" },
    { name: "User2", content: "Nice to meet you!", createdAt: "2025/1/4 12:45" },
    { name: "User3", content: "Great work!", createdAt: "2025/1/4 13:00" },
    { name: "User4", content: "Keep it up!", createdAt: "2025/1/4 14:15" },
  ])

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
              <p class="text-base mt-1">{comment.content}</p>
            </li>
          )}
        </For>
      </ul>
    </>
  );
};

export default CommentList;
