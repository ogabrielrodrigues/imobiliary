import { auth } from "@/actions/auth";
import { getPlan } from "@/actions/plan";
import { AppSidebar } from "@/components/app-sidebar";
import { SiteHeader } from "@/components/site-header";
import { SidebarInset, SidebarProvider } from "@/components/ui/sidebar";
import { PropsWithChildren } from "react";

export default async function DashboardLayout({ children }: PropsWithChildren) {
  const user = await auth()

  const plan = await getPlan()

  return (
    <div className="[--header-height:calc(theme(spacing.14))]">
      <SidebarProvider className="flex flex-col">
        <SiteHeader />
        <div className="flex flex-1">
          <AppSidebar user={user} plan={plan} />
          <SidebarInset className="p-4 sm:p-8 overflow-x-hidden">
            {children}
          </SidebarInset>
        </div>
      </SidebarProvider>
    </div>
  );
}
