export function PhotoCard({ url, description, displayDate, altDate }) {
  return (
    <div className="card card-compact bg-base-100 w-full sm:w-auto sm:max-w-md shadow-xl">
      {/* TODO: See if custom class "min-w-28rem is best approach*/}
      <figure className={"sm:min-w-28rem"}>
        <img
          className="object-contain"
          src={url}
          alt={`Shows ${description}`}
        />
      </figure>
      <div className="card-body">
        <div className="card-actions justify-end">
          <div title={altDate} className="badge badge-outline badge-sm">
            {displayDate}
          </div>
        </div>
        <h2 className="card-title">{description}</h2>
      </div>
    </div>
  );
}
