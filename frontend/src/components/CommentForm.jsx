import { createSignal } from "solid-js";

const CommentForm = ({ onSubmit }) => {
  const [name, setName] = createSignal("");
  const [email, setEmail] = createSignal("");
  const [content, setContent] = createSignal("");

  const handleSubmit = (e) => {
    e.preventDefault();
    onSubmit({ name: name(), content: content(), email: email() });
    // Clean Previous Content
    setContent("");
  };

  return (
    <form onSubmit={handleSubmit}>
      <div class="flex mb-2 space-x-3">
        <input
          type="text"
          placeholder="昵称"
          value={name()}
          onInput={(e) => setName(e.target.value)}
          class="input input-bordered w-full"
          required
        />
        <input
          type="text"
          placeholder="Email"
          value={email()}
          onInput={(e) => setEmail(e.target.value)}
          class="input input-bordered w-full"
          required
        />
      </div>
      <div class="mb-2">
        <div className="label">
          <span className="label-text">内容</span>
          <span className="label-text-alt text-slate-500">悄悄告诉你：什么都不想说也没关系的</span>
        </div>
        <textarea
          placeholder="想说点什么？"
          value={content()}
          onInput={(e) => setContent(e.target.value)}
          class="textarea textarea-bordered textarea-md w-full min-h-36"
          required
        ></textarea>
      </div>
      <div class="flex justify-end items-center">
        <button type="submit" class="ml-auto btn btn-secondary max-w-25">发表评论</button>
      </div>
    </form>
  );
};

export default CommentForm;