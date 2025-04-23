import { auth } from "@/actions/queries/auth";
import { getPlan } from "@/actions/queries/plan/get-plan";
import { AppSidebar } from "@/components/app-sidebar";
import { SiteHeader } from "@/components/site-header";
import { SidebarInset, SidebarProvider } from "@/components/ui/sidebar";
import { PropsWithChildren } from "react";

export default async function DashboardLayout({ children }: PropsWithChildren) {
  const user = await auth()

  const { status, plan } = await getPlan()

  if (status !== 200) {
    return (
      <div className="w-full h-screen flex flex-col items-center justify-center gap-4">
        <h1 className="text-2xl lg:text-7xl font-bold">Erro</h1>

        <p className="text-md lg:text-lg text-center text-muted-foreground">
          Não foi possível carregar o plano. Tente novamenete mais tarde.
        </p>
      </div>
    )
  }

  return (
    <div className="[--header-height:calc(theme(spacing.14))]">
      <SidebarProvider className="flex flex-col">
        <SiteHeader />
        <div className="flex flex-1">
          <AppSidebar user={user} plan={plan!} />
          <SidebarInset className="p-4 sm:p-8 overflow-x-hidden">
            {children}
          </SidebarInset>
        </div>
      </SidebarProvider>
    </div>
  );
}
