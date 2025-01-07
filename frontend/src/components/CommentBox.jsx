import { createSignal, onMount } from "solid-js";
import CommentForm from "./CommentForm";
import CommentList from "./CommentList";



const CommentBox = () => {

  const fetchComments = async () => {
    const response = await fetch(`http://127.0.0.1:3001/api/page/get`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({
        site: "diredocks.github.io",
        path: "/blog/",
      }),
    });
    const data = await response.json();
    setComments(data.Comments);
  }

  const [comments, setComments] = createSignal([]);

  const addComment = (comment) => {
    setComments([...comments(), comment]);
  }

  onMount(fetchComments)

  return (
    <div>
      <CommentForm onSubmit={addComment} />
      <CommentList comments={comments} />
    </div>
  )
};

export default CommentBox;
