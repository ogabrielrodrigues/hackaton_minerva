'use client'

import { feedbackAction } from '@/actions/feedback'
import { useRef } from 'react'
import { useFormStatus } from 'react-dom'
import { useFormState } from 'react-dom'

type Status = 'null' | 'error' | 'success'

const feedback_status: { feedback: Status } = {
  feedback: 'null',
}

function Submit() {
  const { pending } = useFormStatus()

  return (
    <button
      type="submit"
      aria-disabled={pending}
      className="h-10 px-4 bg-primary text-zinc-50 rounded cursor-pointer"
    >
      Enviar
    </button>
  )
}

export function FeedbackForm() {
  const [, feedback_action] = useFormState(feedbackAction, feedback_status)
  const ref = useRef<HTMLFormElement>(null)

  return (
    <form
      ref={ref}
      action={async fd => {
        await feedback_action(fd)
        ref.current?.reset()
      }}
      className="w-full flex flex-col gap-2"
    >
      <div className="flex flex-col gap-1">
        <label htmlFor="name" className="font-medium">
          Novo feedback
        </label>
        <div className="flex gap-2">
          <textarea
            id="content"
            name="content"
            className="w-full resize-y p-2 border border-zinc-200 rounded pl-2 outline-none focus:ring-2 focus:ring-primary transition-all"
            placeholder="Seu feedback"
          />
        </div>
        <div className="w-full flex justify-end">
          <Submit />
        </div>
      </div>
    </form>
  )
}
