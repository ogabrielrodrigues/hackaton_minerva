'use client'

import { useState } from 'react'
import { AnswerModal } from './answer_modal'

interface Modal {
  open: boolean
  feedback_id: string
}

interface Props {
  feedback_id: string
}

export function AnswerButton({ feedback_id }: Props) {
  const [answer_modal, setAnswerModal] = useState<Modal>({
    open: false,
    feedback_id,
  })

  function toggleModal() {
    setAnswerModal({
      open: !answer_modal.open,
      feedback_id,
    })
  }

  return (
    <>
      <button
        onClick={toggleModal}
        className="px-2 py-1 bg-primary text-zinc-50 rounded cursor-pointer"
      >
        Responder
      </button>
      <AnswerModal
        open={answer_modal.open}
        feedback_id={answer_modal.feedback_id}
        toggleModal={toggleModal}
      />
    </>
  )
}
