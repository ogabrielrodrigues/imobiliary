'use server'

import { ContactFormValues } from "@/app/contato/_components/contact-form";
import { createTransport } from 'nodemailer';

export async function sendMail(values: ContactFormValues): Promise<number> {
  const { name, email, message } = values

  const transporter = createTransport({
    from: email,
    host: process.env.SMTP_HOST,
    port: Number(process.env.SMTP_PORT),
    auth: {
      user: process.env.SMTP_USER,
      pass: process.env.SMTP_PASS,
    },
  })

  const { response } = await transporter.sendMail({
    from: email,
    to: process.env.CONTACT_EMAIL,
    subject: `Contato - ${name}`,
    text: message,
  })

  if (response.includes("250")) {
    return 250
  }

  return 500
}