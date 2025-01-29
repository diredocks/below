import "./style/List.css";
import { createSignal } from "solid-js";
import { Index } from "solid-js/web";

const initialComments = [
  {
    id: 1,
    author: "Jane Doe",
    content: "This is a great post!",
    date: "2025-01-29",
  },
  {
    id: 2,
    author: "John Smith",
    content: "I completely agree!",
    date: "2025-01-29",
  },
];

export default function CommentList() {
  const [comments, setComments] = createSignal(initialComments);

  return (
    <section class="comment-list">
      <ul>
        <Index each={comments()}>{(comment) => <CommentItem comment={comment()} />}</Index>
      </ul>
    </section>
  );
}

function CommentItem(props) {
  return (
    <li>
      <article>
        <header>
          <strong>{props.comment.author}</strong>{" "}
          <time datetime={props.comment.date}>{props.comment.date}</time>
        </header>
        <p>{props.comment.content}</p>
      </article>
    </li>
  );
}
