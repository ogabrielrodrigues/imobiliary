import { getManager } from "@/actions/queries/manager/get-manager"
import { listProperties } from "@/actions/queries/property/list-properties"
import { Card, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import type { Metadata } from "next"
import { DashboardHeader } from "./_components/dashboard-header"

export const metadata: Metadata = {
  title: "Dashboard"
}

export default async function DashboardPage() {
  const { manager } = await getManager()

  const dateFormatter = new Intl.DateTimeFormat('pt-BR', { dateStyle: 'short' })

  const { properties: foundProperties } = await listProperties()
  const properties = !foundProperties ? [] : foundProperties

  return (
    <div className="container mx-auto flex flex-col space-y-6">
      <DashboardHeader properties={properties} />
      <div className="text-muted-foreground">
        <div className="flex items-center justify-between font-heading text-lg text-muted">
          <h1>Olá, {manager?.fullname.split(" ")[0]}</h1>
          <span className="text-muted-foreground">{dateFormatter.format(Date.now())}</span>
        </div>
        <p>Aqui você tem visão geral do funcionamento da administração de seus imóveis</p>
      </div>
      <section className="flex flex-col gap-4 mb-4 sm:mb-0">
        <div className="grid sm:grid-cols-2 lg:grid-cols-4 gap-6">
          <Card>
            <CardHeader>
              <CardTitle>Ocorrências</CardTitle>
              <CardDescription>Últimos registros de locação</CardDescription>
            </CardHeader>
          </Card>
          <Card>
            <CardHeader>
              <CardTitle>Anúncios</CardTitle>
              <CardDescription>Imóveis anúnciados na seção pública</CardDescription>
            </CardHeader>
          </Card>
          <Card>
            <CardHeader>
              <CardTitle>Relatórios</CardTitle>
              <CardDescription>Emissão de relatórios de operação</CardDescription>
            </CardHeader>
          </Card>
          <Card>
            <CardHeader>
              <CardTitle>Financeiro</CardTitle>
              <CardDescription>Controle de caixa e extrato</CardDescription>
            </CardHeader>
          </Card>
          <Card>
            <CardHeader>
              <CardTitle>Notificações</CardTitle>
              <CardDescription>Histórico de notificações enviadas/recebidas</CardDescription>
            </CardHeader>
          </Card>
          <Card>
            <CardHeader>
              <CardTitle>Reajuste aluguel</CardTitle>
              <CardDescription>Simule e calcule o reajuste do aluguel pelo FGV/IGP-M</CardDescription>
            </CardHeader>
          </Card>
          <Card>
            <CardHeader>
              <CardTitle>Vistorias</CardTitle>
              <CardDescription>Laudos e fotos da vistorias realizadas</CardDescription>
            </CardHeader>
          </Card>
        </div>
      </section>
    </div>
  )
}