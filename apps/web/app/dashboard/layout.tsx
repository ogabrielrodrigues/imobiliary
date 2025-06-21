import { auth } from "@/actions/queries/auth/auth";
import { getManager } from "@/actions/queries/manager/get-manager";
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

  const { manager, status } = await getManager()
  if (!manager || status != 200) {
    redirect("/login")
  }

  return (
    <div className="[--header-height:calc(theme(spacing.14))]">
      <SidebarProvider className="flex flex-col">
        <div className="flex flex-1 !bg-primary/3">
          <AppSidebar manager={manager} />
          <SidebarInset className="!m-0 w-full lg:!m-3">
            <SiteHeader />
            <section className="px-4">
              {children}
            </section>
          </SidebarInset>
        </div>
      </SidebarProvider>
    </div>
  );
}
