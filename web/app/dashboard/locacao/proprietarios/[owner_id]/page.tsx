import { getOwner } from "@/actions/queries/owner/get-owner"
import { Button } from "@/components/ui/button"
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import { ArrowLeft, EllipsisVertical } from "lucide-react"
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
    <>
      <div className="absolute z-10 w-1/2 h-3 blur-[92px] top-0 left-1/2 -translate-x-1/2 bg-white" />
      <div className="mx-auto max-w-3xl w-full flex flex-col space-y-4">
        <div className="flex justify-between gap-4">
          <div>
            <h1 className="text-2xl font-bold">{owner.fullname}</h1>
            <span className="text-muted-foreground text-sm">{owner.address.mini_address}</span>
          </div>
          <div className="flex flex-col sm:flex-row gap-2">
            <Link href="/dashboard/locacao/proprietarios">
              <Button variant="outline">
                <ArrowLeft className="w-4 h-4" />
                <p className="hidden md:block">Voltar</p>
              </Button>
            </Link>
            <Button>
              <EllipsisVertical className="size-4" />
            </Button>
          </div>
        </div>
        <Card className="bg-zinc-900/20 backdrop-blur-2xl relative z-20 overflow-hidden">
          <CardHeader>
            <CardTitle className="text-lg font-bold">Dados p/ contato</CardTitle>
          </CardHeader>
          <CardContent>
            <div className="flex flex-col gap-2">
              <p className="text-muted-foreground"><strong className="text-primary">Email:</strong> {owner.email}</p>
              <p className="text-muted-foreground"><strong className="text-primary">Telefone:</strong> {owner.cellphone}</p>
              <p className="text-muted-foreground"><strong className="text-primary">Endereço:</strong> {owner.address.mini_address}</p>
            </div>
          </CardContent>
        </Card>
      </div>
    </>
  )
}