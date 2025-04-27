import { Separator } from "@/components/ui/separator"
import { ContactForm } from "./_components/contact-form"

export default function Contato() {
  return (
    <div className="mt-16 flex items-center justify-center">
      <main className="max-w-2xl w-full flex flex-col items-center gap-4 p-4">
        <h1 className="text-3xl font-bold">Contato</h1>
        <p className="text-muted-foreground">
          Para entrar em contato, envie um e-mail para{" "}
          <a
            href="mailto:contato@imobiliary.com"
            className="text-white hover:underline"
          >contato@imobiliary.com</a>.
        </p>

        <div className="w-full my-6 flex items-center justify-center gap-2 text-muted-foreground">
          <Separator className="w-1/3" />
          ou
          <Separator className="w-1/3" />
        </div>

        <ContactForm />
      </main>
    </div>
  )
}