import Box from './Box.jsx';
import List from './List.jsx';
import { createSignal } from "solid-js";

export default function Comment() {
  const [comments, setComments] = createSignal([
    { id: 1, author: "Jane Doe", content: "This is a great post!", date: "2025-01-29" },
    { id: 2, author: "John Smith", content: "I completely agree!", date: "2025-01-29" },
  ]);

  const addComment = (author, content) => {
    setComments([...comments(), { id: comments().length + 1, author, content, date: new Date().toISOString().split("T")[0] }]);
  };

  return (
    <section class="comment-widget">
      <Box onAddComment={addComment} />
      <List comments={comments()} />
    </section>
  );
}
