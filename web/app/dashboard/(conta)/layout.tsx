import { SidebarContent, SidebarMenu, SidebarMenuButton, SidebarMenuItem } from "@/components/ui/sidebar"
import { CreditCard, UserRound } from "lucide-react"
import { headers } from "next/headers"
import { PropsWithChildren } from "react"

const paths = [
  {
    url: "/dashboard/conta",
    icon: UserRound,
    label: "Conta"
  },
  {
    url: "/dashboard/assinatura",
    icon: CreditCard,
    label: "Assinatura"
  }
]

export default async function ContaLayout({ children }: PropsWithChildren) {
  const pathname = (await headers()).get('x-pathname')

  return (
    <div className="[--header-height:calc(theme(spacing.14))] max-w-max mx-auto">
      <SidebarContent className="flex flex-col space-y-5 lg:flex-row lg:space-x-5">
        <SidebarMenu className="flex flex-row lg:flex-col gap-1 w-max">
          {paths.map((path) => (
            <SidebarMenuItem key={path.url}>
              <SidebarMenuButton asChild isActive={pathname === path.url}>
                <a href={path.url} className="flex items-center gap-2">
                  <path.icon />
                  <span>{path.label}</span>
                </a>
              </SidebarMenuButton>
            </SidebarMenuItem>
          ))}
        </SidebarMenu>
        {children}
      </SidebarContent>
    </div>
  )
}