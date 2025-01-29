import "./style/List.css";
import { Index } from "solid-js/web";

export default function CommentList(props) {
  return (
    <section class="comment-list">
      <ul>
        <Index each={props.comments}>{(comment) => <CommentItem comment={comment()} />}</Index>
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
