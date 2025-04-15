import { Tooltip } from "@/components/ui/tooltip";

import { TooltipContent } from "@/components/ui/tooltip";

import { AvatarFallback, AvatarImage } from "@/components/ui/avatar";

import { auth, logout } from "@/actions/auth";
import { Avatar } from "@/components/ui/avatar";
import { Button } from "@/components/ui/button";
import { TooltipTrigger } from "@/components/ui/tooltip";
import { HousePlus, LogOut, Sparkles } from "lucide-react";
import Link from "next/link";

export async function HomeHeader() {
  const user = await auth()

  return (
    <header className="fixed top-0 left-0 z-50 w-full bg-zinc-950/95 backdrop-blur-3xl p-4 flex items-center justify-between gap-4">
      <div className="absolute right-8">
        {user && <div className="flex items-center space-x-4">
          <Tooltip>
            <TooltipTrigger asChild>
              <Link href="/dashboard/conta">
                <Avatar>
                  <AvatarImage src={user?.avatar} className="object-cover" />
                  <AvatarFallback className="bg-sidebar-primary">{user?.fullname.charAt(0)}</AvatarFallback>
                </Avatar>
              </Link>
            </TooltipTrigger>
            <TooltipContent>
              <p>Ver Conta</p>
            </TooltipContent>
          </Tooltip>
          <Button asChild variant="ghost">
            <Link href="/dashboard">
              Dashboard
            </Link>
          </Button>
          <form action={logout}>
            <Button variant="outline" type="submit">
              <LogOut className="size-4" />
              <p className="hidden sm:block">Sair</p>
            </Button>
          </form>
        </div>}
        {!user && <div>
          <Link href="/login" id="login">
            <Button variant="ghost">Login</Button>
          </Link>
          <Link href="/cadastro" id="sign">
            <Button variant="ghost">Cadastro</Button>
          </Link>
        </div>}
      </div>
      <div className="flex items-center gap-2">
        <HousePlus className="size-5" />
        <span className="sr-only">Imobiliary</span>
        <>
          <h1 className="text-lg font-bold hidden sm:block select-none">
            Imobiliary
          </h1>
          {user?.plan.kind === 'pro' && (
            <Sparkles className="size-4 self-start" />
          )}
        </>
      </div>
    </header>
  )
}
