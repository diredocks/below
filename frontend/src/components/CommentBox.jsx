import { createSignal } from "solid-js";
import CommentForm from "./CommentForm";
import CommentList from "./CommentList";

const CommentBox = () => {

  const [comments, setComments] = createSignal([
    { name: "Kytoki", content: "Hello there!", createdAt: "2025/1/4 12:34" },
    { name: "Rynimo", content: "Nice to meet you!", createdAt: "2025/1/4 12:45" },
    { name: "Genimis", content: "Great work!", createdAt: "2025/1/4 13:00" },
    { name: "Terraock", content: "Keep it up!", createdAt: "2025/1/4 14:15" },
  ])

  const addComment = (comment) => {
    setComments([...comments(), comment]);
  }

  return (
    <div>
      <CommentForm onSubmit={addComment} />
      <CommentList comments={comments} />
    </div>
  )
};

export default CommentBox;
