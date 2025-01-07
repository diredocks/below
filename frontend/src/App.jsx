import CommentBox from "./components/CommentBox";

function App() {
  return (
    <div class="flex items-center justify-center h-screen bg-base-200">
      {/* container of comment box as demo page */}
      <div class="max-w-xl bg-base-100 rounded-md border shadow-md p-10"><CommentBox/></div>
    </div>
  );
}

export default App;