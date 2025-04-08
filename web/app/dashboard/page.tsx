import type { Metadata } from "next"

import { ChartAreaInteractive } from "@/components/chart-area-interactive"
import { SectionCards } from "@/components/section-cards"
import { SiteHeader } from "@/components/site-header"

export const metadata: Metadata = {
  title: "Dashboard"
}

export default function DashboardPage() {
  return (
    <>
      <SiteHeader title="Visão Geral" />
      <div className="flex flex-1 flex-col">
        <div className="@container/main flex flex-1 flex-col gap-2">
          <div className="h-full flex flex-col gap-4 py-4 md:gap-6 md:py-6">
            <SectionCards />
            <div className="h-full px-4 lg:px-6">
              <ChartAreaInteractive />
            </div>
          </div>
        </div>
      </div>
    </>
  )
}