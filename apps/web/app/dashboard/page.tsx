import type { Metadata } from "next"
import { DashboardHeader } from "./_components/dashboard-header"

export const metadata: Metadata = {
  title: "Dashboard"
}

export default function DashboardPage() {
  return (
    <div className="container mx-auto flex flex-col space-y-10">
      <DashboardHeader />
    </div>
  )
}