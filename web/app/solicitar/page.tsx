import type { Metadata } from "next"
import { RequestForm } from "@/components/request-form"

export const metadata: Metadata = {
  title: "Solicitar"
}

export default function RequestPage() {
  return (
    <div className="bg-background flex flex-col h-screen w-screen overflow-hidden items-center justify-center gap-6 p-6 md:p-10">
      <RequestForm />
    </div>
  )
}