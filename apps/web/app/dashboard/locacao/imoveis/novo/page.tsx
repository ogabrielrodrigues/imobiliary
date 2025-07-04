import { listOwners } from "@/actions/queries/owner/list-owners"
import { Button } from "@/components/ui/button"
import { ArrowLeft } from "lucide-react"
import { Metadata } from "next"
import Link from "next/link"
import { NewPropertyForm } from "../_components/new-property-form"

export const metadata: Metadata = {
  title: "Imóveis - Novo",
}

export default async function NewPropertyPage() {
  const { status: owner_status, owners: found } = await listOwners()

  if (owner_status != 200) {
    return (
      <div className="w-full flex justify-center">
        <p className="font-medium text-muted">Erro ao carregar proprietários</p>
      </div>
    )
  }

  const owners = !found ? [] : found

  return (
    <div className="container mx-auto xl:max-w-xl flex flex-col space-y-4">

      <div className="flex items-center justify-between">
        <h1 className="text-2xl font-bold">Novo Imóvel</h1>
        <Link href="/dashboard/locacao/imoveis">
          <Button variant="outline">
            <ArrowLeft className="w-4 h-4" />
            <p className="hidden md:block">Voltar</p>
          </Button>
        </Link>
      </div>
      <NewPropertyForm owners={owners} />  {/* Fazer melhoria de busca de endereço ...*/}
    </div >
  )
}
