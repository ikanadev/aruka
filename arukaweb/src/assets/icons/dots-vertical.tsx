import type { SVGProps } from "react";

export function DotsVerticalIcon(props: SVGProps<SVGSVGElement>) {
  return (
    <svg xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" viewBox="0 0 24 24" {...props}>
      <title>dots-vertical</title>
      {/* Icon from Mono Icons by Mono - https://github.com/mono-company/mono-icons/blob/master/LICENSE.md */}
      <path
        fill="currentColor"
        d="M12 8a2 2 0 1 1 0-4 2 2 0 0 1 0 4m0 6a2 2 0 1 1 0-4 2 2 0 0 1 0 4m0 6a2 2 0 1 1 0-4 2 2 0 0 1 0 4"
      />
    </svg>
  );
}
