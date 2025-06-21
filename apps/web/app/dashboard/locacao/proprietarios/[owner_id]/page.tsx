import { getOwner } from "@/actions/queries/owner/get-owner"
import { Button } from "@/components/ui/button"
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import { ArrowLeft } from "lucide-react"
import { Metadata } from "next"
import Link from "next/link"
import { notFound } from "next/navigation"

type OwnerDetailsPageParams = {
  params: Promise<{ owner_id: string }>
}

export const metadata: Metadata = {
  title: "Perfil do Proprietário",
}

export default async function OwnerDetailsPage({ params }: OwnerDetailsPageParams) {
  const { owner_id } = await params
  const { status, owner } = await getOwner(owner_id)

  if (!owner || status !== 200) notFound()

  return (
    <div className="mx-auto max-w-3xl w-full flex flex-col space-y-4">
      <div className="flex justify-between gap-4">
        <div>
          <h1 className="text-2xl font-bold">{owner.fullname}</h1>
          <span className="text-muted-foreground text-sm">{owner.address.address}</span>
        </div>
        <div className="flex flex-col sm:flex-row gap-2">
          <Link href="/dashboard/locacao/proprietarios">
            <Button variant="outline">
              <ArrowLeft className="w-4 h-4" />
              <p className="hidden md:block">Voltar</p>
            </Button>
          </Link>
        </div>
      </div>
      <Card className="backdrop-blur-2xl relative z-20 overflow-hidden">
        <CardHeader>
          <CardTitle className="text-lg font-bold">Dados p/ contato</CardTitle>
        </CardHeader>
        <CardContent>
          <div className="flex flex-col gap-2">
            <p className="text-muted-foreground"><strong className="text-muted">Email:</strong> {owner.email}</p>
            <p className="text-muted-foreground"><strong className="text-muted">Telefone:</strong> {owner.phone}</p>
            <p className="text-muted-foreground"><strong className="text-muted">Endereço:</strong> {owner.address.address}</p>
          </div>
        </CardContent>
      </Card>
    </div>
  )
}