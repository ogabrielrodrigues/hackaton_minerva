'use client'

import { loginAction } from '@/actions/login'
import { useFormStatus } from 'react-dom'
import { useFormState } from 'react-dom'

const login_status = {
  success: false,
}

function Submit() {
  const { pending } = useFormStatus()

  return (
    <button
      type="submit"
      aria-disabled={pending}
      className="h-10 bg-primary text-zinc-50 rounded cursor-pointer"
    >
      Entrar
    </button>
  )
}

export function LoginForm() {
  const [, login_action] = useFormState(loginAction, login_status)

  return (
    <form action={login_action} className="flex flex-col gap-2">
      <div className="flex flex-col gap-1">
        <label htmlFor="email" className="font-medium">
          E-mail
        </label>
        <input
          id="email"
          name="email"
          type="text"
          // autoComplete="off"
          className="h-10 border border-zinc-200 rounded pl-2 outline-none focus:ring-2 focus:ring-primary transition-all"
          placeholder="Digite seu e-mail"
        />
      </div>
      <div className="flex flex-col gap-1">
        <label htmlFor="password" className="font-medium">
          Senha
        </label>
        <input
          id="password"
          name="password"
          type="password"
          className="h-10 border border-zinc-200 rounded pl-2 outline-none focus:ring-2 focus:ring-primary transition-all"
          placeholder="Digite sua senha"
        />
      </div>
      <Submit />
    </form>
  )
}
