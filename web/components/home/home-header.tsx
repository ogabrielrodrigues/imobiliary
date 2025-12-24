import {
	IconCirclePlus,
	IconDoorExit,
	IconHomePlus,
	IconLogin2,
} from "@tabler/icons-react"
import Link from "next/link"
import { logout } from "@/actions/mutations/auth/logout"
import { getManager } from "@/actions/queries/manager/get-manager"
import { Button } from "@/components/ui/button"

export async function HomeHeader() {
	const { manager } = await getManager()

	return (
		<header className="w-screen flex items-center justify-between p-4 border-b border-border">
			<div className="flex items-center gap-2 text-muted">
				<IconHomePlus className="size-5" />
				<strong className="text-xl select-none">imobiliary</strong>
			</div>

			{manager ? (
				<form action={logout} className="flex gap-4 items-center">
					<Link className="text-sm font-medium" href="/dashboard">
						{manager?.fullname.split(" ")[0]}
					</Link>

					<Button type="submit" variant="outline">
						<IconDoorExit className="size-4" />
						<p className="hidden sm:block">Sair</p>
					</Button>
				</form>
			) : (
				<div className="flex items-center gap-2">
					<Link href="/login">
						<Button variant="outline">
							<IconLogin2 />
							<span className="hidden sm:block">Login</span>
						</Button>
					</Link>

					<Link href="/cadastro">
						<Button>
							<IconCirclePlus />
							<span className="hidden sm:block">Cadastro</span>
						</Button>
					</Link>
				</div>
			)}
		</header>
	)
}
