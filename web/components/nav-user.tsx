"use client"

import { useRouter } from 'next/navigation'

import { DoorOpen } from "lucide-react"

import { SidebarMenu, SidebarMenuItem } from "@/components/ui/sidebar"
import { Button } from "./ui/button"
import { Tooltip, TooltipContent, TooltipTrigger } from "./ui/tooltip"
import { NavUserProps } from "@/types/nav-items"

export function NavUser({ user }: NavUserProps) {
  const navigate = useRouter()

  function Exit() {
    window.localStorage.removeItem("imobiliary-auth")
    navigate.push("/")
  }

  return (
    <SidebarMenu>
      <SidebarMenuItem className="flex gap-2">
        <div className="grid flex-1 text-left text-sm leading-tight">
          <span className="truncate font-medium">{user.name}</span>
          <span className="text-muted-foreground truncate text-xs">
            {user.email}
          </span>
        </div>
        <Tooltip>
          <TooltipTrigger asChild>
            <Button
              variant="outline"
              onClick={Exit}
            >
              <DoorOpen className="ml-auto size-4" />
            </Button>
          </TooltipTrigger>
          <TooltipContent>
            Sair
          </TooltipContent>
        </Tooltip>
      </SidebarMenuItem>
    </SidebarMenu>
  )
}
