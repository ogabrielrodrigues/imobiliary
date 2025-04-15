"use client"

import {
  CreditCard,
  EllipsisVertical,
  LogOut,
  Sparkles,
  UserRound
} from "lucide-react"

import {
  Avatar,
  AvatarFallback,
  AvatarImage
} from "@/components/ui/avatar"
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

import { User } from "@/types/user"

import { logout } from "@/actions/auth"
import { ProPlanDialog } from "./pro-plan-dialog"

type NavUserProps = {
  user: User | undefined
}

export function NavUser({ user }: NavUserProps) {
  const { isMobile } = useSidebar()

  return (
    <SidebarMenu>
      <SidebarMenuItem>
        <DropdownMenu>
          <DropdownMenuTrigger asChild>
            <SidebarMenuButton
              size="lg"
              className="data-[state=open]:bg-sidebar-accent data-[state=open]:text-sidebar-accent-foreground"
            >
              <Avatar className="h-8 w-8 rounded-lg">
                {user?.avatar ? <AvatarImage src={user?.avatar} className="object-cover" />
                  : <AvatarFallback className="bg-sidebar-primary animate-pulse" />}
              </Avatar>
              <div className="grid flex-1 text-left text-sm leading-tight">
                <span className="truncate font-medium">{user!.fullname}</span>
                <span className="truncate text-xs text-muted-foreground">{user!.email}</span>
              </div>
              <EllipsisVertical className="ml-auto size-4" />
            </SidebarMenuButton>
          </DropdownMenuTrigger>
          <DropdownMenuContent
            className="w-(--radix-dropdown-menu-trigger-width) min-w-56 rounded-lg"
            side={isMobile ? "bottom" : "right"}
            align="end"
            sideOffset={4}
          >
            <DropdownMenuLabel className="p-0 font-normal">
              <div className="flex items-center gap-2 px-1 py-1.5 text-left text-sm">
                <Avatar className="h-8 w-8 rounded-lg">
                  {user?.avatar ? <AvatarImage src={user?.avatar} className="object-cover" />
                    : <AvatarFallback className="bg-sidebar-primary animate-pulse" />}
                </Avatar>
                <div className="grid flex-1 text-left text-sm leading-tight">
                  <span className="truncate font-medium">{user!.fullname}</span>
                  <span className="truncate text-xs text-muted-foreground">{user!.email}</span>
                </div>
              </div>
            </DropdownMenuLabel>
            <DropdownMenuSeparator />
            {user?.plan.kind === "free" && (
              <>
                <DropdownMenuGroup>
                  <ProPlanDialog>
                    <DropdownMenuItem onSelect={(e) => e.preventDefault()}>
                      <Sparkles />
                      Atualizar para o PRO
                    </DropdownMenuItem>
                  </ProPlanDialog>
                </DropdownMenuGroup>
                <DropdownMenuSeparator />
              </>
            )}
            <DropdownMenuGroup>
              <DropdownMenuItem asChild>
                <a href="/dashboard/conta">
                  <UserRound />
                  Conta
                </a>
              </DropdownMenuItem>
              <DropdownMenuItem asChild>
                <a href="/dashboard/assinatura">
                  <CreditCard />
                  <span>Assinatura</span>
                </a>
              </DropdownMenuItem>
            </DropdownMenuGroup>
            <DropdownMenuSeparator />
            <form action={logout}>
              <DropdownMenuItem asChild>
                <button type="submit" className="w-full">
                  <LogOut />
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
