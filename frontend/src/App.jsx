import CommentBox from "./components/CommentBox";
import CtaSection from "./components/CtaSection";

function App() {
  return (
    <div class="flex flex-col items-center justify-center h-full min-h-screen bg-base-200">
      {/*<div class="cta-section p-20 m-20 pb-10 mb-10 text-balance"><CtaSection /></div>/}
      {/* container of comment box as demo page */}
      <div class="comment-box max-w-3xl bg-base-100 rounded-md border shadow-md p-10 m-10"><CommentBox /></div>
    </div >
  );
}

export default App;
