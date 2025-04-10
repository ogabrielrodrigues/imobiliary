import { auth } from "@/actions/auth"
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar"
import { Card, CardContent, CardDescription, CardFooter, CardHeader } from "@/components/ui/card"
import { Metadata } from "next"

export const metadata: Metadata = {
  title: "Conta",
  description: "Conta"
}

export default async function AccountPage() {
  const user = await auth()

  return (
    <div className="relative">
      <div className="hidden sm:block absolute z-10 bg-zinc-50 w-20 h-20 blur-[96px] translate-y-1/2 -translate-x-1/2 top-1/2 left-1/2" />
      <div className="max-w-4xl mx-auto space-y-6">
        <div className="flex justify-between items-center">
          <h1 className="text-3xl font-bold">Conta</h1>
        </div>

        <Card className="z-20 bg-zinc-900/20 backdrop-blur-2xl relative overflow-hidden">
          <CardContent className="flex flex-col items-center sm:flex-row space-x-6 space-y-4">
            <Avatar className="w-20 h-20 text-2xl font-semibold">
              {
                user?.avatar ?
                  <AvatarImage src={user.avatar} /> :
                  <AvatarFallback className="bg-sidebar-primary">{user?.fullname?.charAt(0)}</AvatarFallback>
              }
            </Avatar>
            <div className="flex flex-col items-center sm:items-start">
              <h2 className="text-xl font-semibold">{user?.fullname}</h2>
              <p className="text-muted-foreground">{user?.email}</p>
              <span className="inline-block bg-primary/10 text-primary text-xs px-2 py-1 rounded mt-1">Administrador</span>
            </div>
          </CardContent>
        </Card>

        <Card className="!gap-2 z-20 bg-zinc-900/20 backdrop-blur-2xl">
          <CardHeader>
            <CardDescription>Nome completo</CardDescription>
          </CardHeader>
          <CardContent className="text-sm sm:text-base">{user?.fullname}</CardContent>
        </Card>

        <Card className="!gap-2 z-20 bg-zinc-900/20 backdrop-blur-2xl">
          <CardHeader>
            <CardDescription>E-mail</CardDescription>
          </CardHeader>
          <CardContent className="text-sm sm:text-base">{user?.email}</CardContent>
          <CardFooter className="text-xs text-muted-foreground">
            Caso deseje alterar seu e-mail entre em contato com o suporte.
          </CardFooter>
        </Card>

        <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
          <Card className="!gap-2 z-20 bg-zinc-900/20 backdrop-blur-2xl">
            <CardHeader>
              <CardDescription>Telefone</CardDescription>
            </CardHeader>
            <CardContent className="text-sm sm:text-base">{user?.cellphone}</CardContent>
          </Card>
          <Card className="!gap-2 z-20 bg-zinc-900/20 backdrop-blur-2xl">
            <CardHeader>
              <CardDescription>CRECI</CardDescription>
            </CardHeader>
            <CardContent className="text-sm sm:text-base">{user?.creci_id}</CardContent>
          </Card>
        </div>
      </div>
    </div>
  )
}