export function LoadingSpinner() {
  return (
    <div role="status">
      <div
        aria-hidden="true"
        className="radial-progress
        animate-spin"
        style="
          --size:2rem;
          --value:70;
          animation:spin 1s cubic-bezier(.8,-.5,.2,1.4) infinite
      "
      />
      <span className="sr-only">Loading...</span>
    </div>
  );
}
