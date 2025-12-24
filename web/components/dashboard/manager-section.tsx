"use client"

import { IconDoorExit, IconDotsVertical, IconUser } from "@tabler/icons-react"
import { logout } from "@/actions/mutations/auth/logout"
import {
	DropdownMenu,
	DropdownMenuContent,
	DropdownMenuGroup,
	DropdownMenuItem,
	DropdownMenuLabel,
	DropdownMenuSeparator,
	DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu"
import {
	SidebarMenu,
	SidebarMenuButton,
	SidebarMenuItem,
	useSidebar,
} from "@/components/ui/sidebar"
import type { Manager } from "@/types/manager"

interface ManagerSectionProps {
	manager?: Manager
}

export function ManagerSection({ manager }: ManagerSectionProps) {
	const { isMobile } = useSidebar()

	return (
		<SidebarMenu>
			<SidebarMenuItem>
				<DropdownMenu>
					<DropdownMenuTrigger asChild>
						<SidebarMenuButton
							className="data-[state=open]:bg-sidebar-accent data-[state=open]:text-sidebar-accent-foreground"
							size="lg"
						>
							<div className="grid flex-1 text-left text-sm leading-tight">
								<span className="truncate font-medium">
									{manager!.fullname}
								</span>
								<span className="truncate text-xs text-muted-foreground">
									{manager!.email}
								</span>
							</div>
							<IconDotsVertical className="ml-auto size-4" />
						</SidebarMenuButton>
					</DropdownMenuTrigger>

					<DropdownMenuContent
						align="end"
						className="w-(--radix-dropdown-menu-trigger-width) min-w-56 rounded-lg"
						side={isMobile ? "bottom" : "right"}
						sideOffset={4}
					>
						<DropdownMenuLabel className="p-0 font-normal">
							<div className="flex items-center gap-2 px-1 py-1.5 text-left text-sm">
								<div className="grid flex-1 text-left text-sm leading-tight">
									<span className="truncate font-medium">
										{manager!.fullname}
									</span>
									<span className="truncate text-xs text-muted-foreground">
										{manager!.email}
									</span>
								</div>
							</div>
						</DropdownMenuLabel>

						<DropdownMenuSeparator />

						<DropdownMenuGroup>
							<DropdownMenuItem asChild>
								<a className="text-muted-foreground" href="/dashboard/conta">
									<IconUser />
									Conta
								</a>
							</DropdownMenuItem>
						</DropdownMenuGroup>

						<DropdownMenuSeparator />

						<form action={logout}>
							<DropdownMenuItem asChild>
								<button className="w-full text-muted-foreground" type="submit">
									<IconDoorExit />
									Sair
								</button>
							</DropdownMenuItem>
						</form>
					</DropdownMenuContent>
				</DropdownMenu>
			</SidebarMenuItem>
		</SidebarMenu>
	)
}
