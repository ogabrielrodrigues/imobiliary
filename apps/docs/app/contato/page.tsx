import { Metadata } from "next";

export const metadata: Metadata = {
  title: "Contato"
};

export default async function ContactPage() {
  return (
    <div className="mt-16 flex items-center justify-center">
      <main className="max-w-2xl w-full flex flex-col items-center gap-4 p-4">
        <h1 className="text-3xl font-bold">Contato</h1>
        <p className="text-muted-foreground">
          Para entrar em contato, envie um e-mail para{" "}
          <a
            href="mailto:contato@imobiliary.com"
            className="text-primary hover:underline"
          >contato@imobiliary.com</a>.
        </p>
      </main>
    </div>
  )
}