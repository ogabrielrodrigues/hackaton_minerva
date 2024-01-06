'use client'

import { answerAction } from '@/actions/answer'
import { useRef } from 'react'
import { useFormState, useFormStatus } from 'react-dom'

type Answer = 'null' | 'error' | 'success'

const answer_content: { answer: Answer } = {
  answer: 'null',
}

function SubmitButton() {
  const { pending } = useFormStatus()

  return (
    <button
      type="submit"
      aria-disabled={pending}
      className="h-10 px-2 bg-primary text-zinc-50 rounded cursor-pointer"
    >
      Responder
    </button>
  )
}

interface Props {
  feedback_id: string
  toggleModal(): void
}

export function AnswerForm({ feedback_id, toggleModal }: Props) {
  const [answer, answer_action] = useFormState(answerAction, answer_content)
  const ref = useRef<HTMLFormElement>(null)

  return (
    <form
      ref={ref}
      action={async fd => {
        fd.append('feedback_id', feedback_id)
        await answer_action(fd)
        if (answer.answer == 'success') toggleModal()
      }}
      className="w-full flex flex-col gap-2 "
    >
      <div className="flex flex-col gap-1">
        <label htmlFor="sector" className="font-medium">
          Resposta
        </label>
        <textarea
          id="answer"
          name="answer"
          className="border h-64 border-zinc-200 rounded p-2 resize-none outline-none focus:ring-2 focus:ring-primary transition-all"
          placeholder="Digite sua resposta"
        />
        <div className="w-full flex justify-between">
          <p className="text-primary font-medium text-center mt-2">
            {answer.answer == 'error' && 'Mínimo de caracteres: 5'}
          </p>
          <SubmitButton />
        </div>
      </div>
    </form>
  )
}
