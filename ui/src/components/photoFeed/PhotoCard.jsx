import { constants } from "../../lib/constants.js";

export function PhotoCard({
  url,
  description,
  displayDate,
  altDate,
  isNew,
  tags,
}) {
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
          <div title={altDate} className="badge bae-outline badge-sm">
            {displayDate}
          </div>
        </div>
        <h2 className="card-title">
          {description}
          {isNew && <span className="badge badge-secondary">New</span>}
        </h2>
        <div className="card-actions">
          {tags.map((tag, index) => {
            return (
              <a
                className="link link-hover"
                key={index}
                href={constants.ROUTES.getTagRoute(tag.name)}
              >
                #{tag.name}
              </a>
            );
          })}
        </div>
      </div>
    </div>
  );
}
