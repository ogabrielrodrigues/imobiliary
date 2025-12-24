import { redirect } from "next/navigation"
import type { PropsWithChildren } from "react"
import { auth } from "@/actions/queries/auth/auth"
import { getManager } from "@/actions/queries/manager/get-manager"
import { DashboardHeader } from "@/components/dashboard/dashboard-header"
import { DashboardSidebar } from "@/components/dashboard/dashboard-sidebar"
import { SidebarInset, SidebarProvider } from "@/components/ui/sidebar"

export default async function DashboardLayout({ children }: PropsWithChildren) {
	const isAuthenticated = await auth()
	if (!isAuthenticated) {
		redirect("/login")
	}

	const { manager } = await getManager()
	if (!manager) {
		redirect("/login")
	}

	return (
		<div className="[--header-height:calc(theme(spacing.14))]">
			<SidebarProvider className="flex flex-col">
				<div className="flex flex-1 !bg-primary/3">
					<DashboardSidebar manager={manager} />

					<SidebarInset className="!m-0 w-full lg:!m-3">
						<DashboardHeader />

						<section className="px-4">{children}</section>
					</SidebarInset>
				</div>
			</SidebarProvider>
		</div>
	)
}
