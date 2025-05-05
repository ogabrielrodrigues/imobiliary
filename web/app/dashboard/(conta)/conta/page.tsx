import { auth } from "@/actions/queries/auth"
import { getUser } from "@/actions/queries/user/get-user"
import { Card, CardContent, CardDescription, CardFooter, CardHeader } from "@/components/ui/card"
import { Metadata } from "next"
import { redirect } from "next/navigation"
import { AvatarForm } from "./_components/avatar-form"

export const metadata: Metadata = {
  title: "Conta",
  description: "Conta"
}

export default async function AccountPage() {
  const auth_id = await auth()
  if (!auth_id) {
    redirect("/login")
  }

  const { user, status } = await getUser(auth_id)
  if (!user || status != 200) {
    redirect("/login")
  }

  return (
    <div className="lg:w-xl mx-auto space-y-6">
      <div className="flex justify-between items-center">
        <h1 className="text-3xl font-bold">Conta</h1>
      </div>

      <Card className="z-20 bg-zinc-900/20 backdrop-blur-2xl relative overflow-hidden">
        <CardContent className="flex flex-col items-center sm:flex-row space-x-6 space-y-4">
          <AvatarForm user={user} />
          <div className="flex flex-col items-center sm:items-start">
            <h2 className="text-lg md:text-xl font-semibold">{user?.fullname}</h2>
            <p className="text-xs md:text-base text-muted-foreground">{user?.email}</p>
            <span className="inline-block bg-primary/10 text-primary text-xs px-2 py-1 rounded mt-1">Administrador</span>
          </div>
        </CardContent>
      </Card>

      <Card className="!gap-0 z-20 bg-zinc-900/20 backdrop-blur-2xl">
        <CardHeader>
          <CardDescription>Nome completo</CardDescription>
        </CardHeader>
        <CardContent className="text-sm sm:text-base">{user?.fullname}</CardContent>
      </Card>

      <Card className="!gap-0 z-20 bg-zinc-900/20 backdrop-blur-2xl">
        <CardHeader>
          <CardDescription>E-mail</CardDescription>
        </CardHeader>
        <CardContent className="text-xs sm:text-base">{user?.email}</CardContent>
        <CardFooter className="mt-1 text-xs text-muted-foreground">
          Caso deseje alterar seu e-mail entre em contato com o suporte.
        </CardFooter>
      </Card>

      <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
        <Card className="!gap-0 z-20 bg-zinc-900/20 backdrop-blur-2xl">
          <CardHeader>
            <CardDescription>Telefone</CardDescription>
          </CardHeader>
          <CardContent className="text-sm sm:text-base">{user?.cellphone}</CardContent>
        </Card>
        <Card className="!gap-0 z-20 bg-zinc-900/20 backdrop-blur-2xl">
          <CardHeader>
            <CardDescription>CRECI</CardDescription>
          </CardHeader>
          <CardContent className="text-sm sm:text-base">{user?.creci_id}</CardContent>
        </Card>
      </div>
    </div>
  )
}