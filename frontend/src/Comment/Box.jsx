import "./style/Box.css";

export default function Box(props) {
  let contentInput, authorInput;

  const handleSubmit = () => {
    if (authorInput.value.trim() && contentInput.value.trim()) {
      // TODO: Show a warning when not inputed anything, maybe in css
      props.onAddComment(authorInput.value.trim(), contentInput.value.trim());
      contentInput.value = "";
    }
  };

  return (
    <section class="comment-box">
      <label>Say whatever...</label>
      <textarea
        ref={(el) => (contentInput = el)}
        id="content"
        rows="5" cols="33"
        placeholder="It was a dark and stormy night...">
      </textarea>
      <input
        ref={(el) => (authorInput = el)}
        type="text"
        id="username" name="username"
        placeholder="Your name" />
      <input
        type="text"
        id="email" name="email"
        pattern="email"
        placeholder="Your Email (optional)" />
      <button
        onClick={handleSubmit}
      >
        Submit
      </button>
    </section>
  );
}
