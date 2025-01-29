import "./style/Box.css";

export default function Box() {
  return (
    <section class="comment-box">
      <label>Say whatever...</label>
      <textarea
        id="content"
        rows="5" cols="33"
        placeholder="It was a dark and stormy night...">
      </textarea>
      <input
        type="text"
        id="username" name="username"
        placeholder="Your name" />
      <input
        type="text"
        id="email" name="email"
        pattern="email"
        placeholder="Your Email (optional)" />
      <button>
        Submit
      </button>
    </section>
  );
}
