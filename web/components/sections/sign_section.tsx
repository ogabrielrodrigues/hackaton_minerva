import Link from 'next/link'
import { SignForm } from '../forms/sign_form'

export function SignSection() {
  return (
    <>
      <div className="h-3/5 flex flex-col justify-center gap-4">
        <h1 className="text-4xl font-bold">Cadastro</h1>

        <SignForm />
      </div>
      <div className="h-1/5 flex items-end justify-center">
        <Link href="/" className="text-primary font-medium">
          Entrar
        </Link>
      </div>
    </>
  )
}
