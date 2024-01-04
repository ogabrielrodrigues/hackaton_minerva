import Link from 'next/link'
import { LoginForm } from '../forms/login_form'

export function LoginSection() {
  return (
    <>
      <div className="h-1/2 flex flex-col justify-center gap-4">
        <h1 className="text-4xl font-bold">Entrar</h1>

        <LoginForm />
      </div>
      <div className="h-1/3 flex items-end justify-center">
        <Link href="/sign" className="text-primary font-medium">
          Cadastro
        </Link>
      </div>
    </>
  )
}
