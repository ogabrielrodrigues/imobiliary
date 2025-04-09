'use client'

import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar"
import { User } from "@/types/user"

export function UserLink({ user }: { user: User }) {
  return (
    <a href="/dashboard/conta" className="flex items-center gap-2">
      <Avatar className="size-8 rounded-lg">
        {
          user.avatar ?
            <AvatarImage src={user.avatar} /> :
            <AvatarFallback className="bg-sidebar-primary">{user.fullname.charAt(0)}</AvatarFallback>
        }
      </Avatar>

      <div className="grid flex-1 text-left text-sm leading-tight">
        <span className="truncate font-medium">{user.fullname}</span>
        <span className="text-muted-foreground truncate text-xs">{user.email}</span>
      </div>
    </a>
  )
} 