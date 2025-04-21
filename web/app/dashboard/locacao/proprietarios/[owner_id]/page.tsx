import { getOwner } from "@/actions/owner"
import { Button } from "@/components/ui/button"
import { Card, CardContent } from "@/components/ui/card"
import { ArrowLeft, EllipsisVertical } from "lucide-react"
import Link from "next/link"
import { notFound } from "next/navigation"

type OwnerDetailsPageParams = {
  params: Promise<{ owner_id: string }>
}

export default async function OwnerDetailsPage({ params }: OwnerDetailsPageParams) {
  const { owner_id } = await params
  const { status, owner } = await getOwner(owner_id)

  if (status !== 200) notFound()

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
          <CardContent>
            <div className="flex flex-col gap-2">
              <h2 className="text-lg font-bold">Dados p/ contato</h2>
              <p><strong>Email:</strong> {owner.email}</p>
              <p><strong>Telefone:</strong> {owner.cellphone}</p>
            </div>
          </CardContent>
        </Card>
        <div className="flex space-x-2">
          {/* <StatusBadge status={property.status} />
          <Badge className="h-8" variant="outline">{property.kind}</Badge> */}
        </div>
      </div>
    </>
  )
}