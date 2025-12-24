import { format } from "date-fns"
import type { Metadata } from "next"
import { getManager } from "@/actions/queries/manager/get-manager"
import { listProperties } from "@/actions/queries/property/list-properties"
import { DashboardHeaderSection } from "@/components/dashboard/dashboard-header-section"
import {
	Card,
	CardDescription,
	CardHeader,
	CardTitle,
} from "@/components/ui/card"

export const metadata: Metadata = {
	title: "Dashboard",
}

export default async function DashboardPage() {
	const { manager } = await getManager()

	const { properties: foundProperties } = await listProperties()
	const properties = !foundProperties ? [] : foundProperties

	return (
		<div className="w-full flex flex-col space-y-6">
			<DashboardHeaderSection properties={properties} />

			<div className="text-muted-foreground">
				<div className="flex items-center justify-between font-heading text-lg text-muted">
					<h1>Olá, {manager?.fullname.split(" ")[0]}</h1>
					<span className="text-muted-foreground">
						{format(Date.now(), "dd/MM/yyyy")}
					</span>
				</div>

				<p>
					Aqui você tem visão geral do funcionamento da administração de seus
					imóveis
				</p>
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

							<CardDescription>
								Imóveis anúnciados na seção pública
							</CardDescription>
						</CardHeader>
					</Card>

					<Card>
						<CardHeader>
							<CardTitle>Relatórios</CardTitle>

							<CardDescription>
								Emissão de relatórios de operação
							</CardDescription>
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

							<CardDescription>
								Histórico de notificações enviadas/recebidas
							</CardDescription>
						</CardHeader>
					</Card>

					<Card>
						<CardHeader>
							<CardTitle>Reajuste aluguel</CardTitle>

							<CardDescription>
								Simule e calcule o reajuste do aluguel pelo FGV/IGP-M
							</CardDescription>
						</CardHeader>
					</Card>

					<Card>
						<CardHeader>
							<CardTitle>Vistorias</CardTitle>

							<CardDescription>
								Laudos e fotos da vistorias realizadas
							</CardDescription>
						</CardHeader>
					</Card>
				</div>
			</section>
		</div>
	)
}
