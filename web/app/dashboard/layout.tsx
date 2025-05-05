import { auth } from "@/actions/queries/auth";
import { getUser } from "@/actions/queries/user/get-user";
import { AppSidebar } from "@/components/app-sidebar";
import { SiteHeader } from "@/components/site-header";
import { SidebarInset, SidebarProvider } from "@/components/ui/sidebar";
import { redirect } from "next/navigation";
import { PropsWithChildren } from "react";

export default async function DashboardLayout({ children }: PropsWithChildren) {
  const auth_id = await auth()
  if (!auth_id) {
    redirect("/login")
  }

  const { user, status } = await getUser(auth_id)
  if (!user || status != 200) {
    redirect("/login")
  }

  return (
    <div className="[--header-height:calc(theme(spacing.14))]">
      <SidebarProvider className="flex flex-col">
        <SiteHeader />
        <div className="flex flex-1">
          <AppSidebar user={user} />
          <SidebarInset className="p-4 sm:p-8 overflow-x-hidden">
            {children}
          </SidebarInset>
        </div>
      </SidebarProvider>
    </div>
  );
}
