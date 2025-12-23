import { IconHomePlus } from "@tabler/icons-react"
import Link from "next/link"

interface AuthHeaderProps {
	description: string
	href: string
	label: string
}

export function AuthHeader({ description, href, label }: AuthHeaderProps) {
	return (
		<div className="flex flex-col items-center gap-2">
			<Link className="flex gap-2 items-center text-muted" href="/">
				<IconHomePlus className="size-6" />
				<strong className="text-2xl select-none">imobiliary</strong>
			</Link>

			<div className="text-center text-sm text-muted-foreground">
				<span className="mr-1">{description}</span>

				<a className="underline underline-offset-2 text-primary" href={href}>
					{label}
				</a>
			</div>
		</div>
	)
}
