INSERT INTO "tb_employees" (name, email, sector, unit, administrator, password) 
VALUES
('Admin', 'admin@admin.minerva.com', 'Administração', 'Barretos', true, MD5("admin1")),
('John Doe', 'john@ti.minerva.com', 'Informática', 'Barretos', false, MD5("john01")),
('Ana Silva', 'ana.silva@example.com', 'Administração', 'São Paulo', true, MD5('Lp7sK9')),
('Bruno Oliveira', 'bruno.oliveira@example.com', 'Informática', 'Rio de Janeiro', false, MD5('aR3dY8')),
('Carla Santos', 'carla.santos@example.com', 'Mecânica', 'Belo Horizonte', false, MD5('5qXbZr')),
('Daniel Costa', 'daniel.costa@example.com', 'Produção', 'Brasília', false, MD5('pJ2sN6')),
('Eduarda Pereira', 'eduarda.pereira@example.com', 'Frigorífico', 'Salvador', false, MD5('fG8hT3')),
('Fernando Lima', 'fernando.lima@example.com', 'Administração', 'Porto Alegre', true, MD5('9aUzE1')),
('Gabriela Rocha', 'gabriela.rocha@example.com', 'Informática', 'Fortaleza', false, MD5('cH4mY7')),
('Henrique Almeida', 'henrique.almeida@example.com', 'Mecânica', 'Recife', false, MD5('2wRvL5')),
('Isabela Santos', 'isabela.santos@example.com', 'Produção', 'Manaus', false, MD5('xQ6kP3')),
('João Souza', 'joao.souza@example.com', 'Frigorífico', 'Curitiba', false, MD5('7dGjR1')),
('Laura Martins', 'laura.martins@example.com', 'Administração', 'Belém', true, MD5('oI9tM4')),
('Lucas Oliveira', 'lucas.oliveira@example.com', 'Informática', 'Vitória', false, MD5('zK5pQ7')),
('Mariana Silva', 'mariana.silva@example.com', 'Mecânica', 'Natal', false, MD5('1vUyT6')),
('Nathan Pereira', 'nathan.pereira@example.com', 'Produção', 'São Luís', false, MD5('bA8sN3')),
('Olívia Lima', 'olivia.lima@example.com', 'Frigorífico', 'Campo Grande', false, MD5('lJ3wR9')),
('Pedro Santos', 'pedro.santos@example.com', 'Administração', 'João Pessoa', true, MD5('rX7hY2')),
('Quitéria Almeida', 'quiteria.almeida@example.com', 'Informática', 'Aracaju', false, MD5('3qZtM6')),
('Rafael Costa', 'rafael.costa@example.com', 'Mecânica', 'Teresina', false, MD5('6eUoP4')),
('Sílvia Oliveira', 'silvia.oliveira@example.com', 'Produção', 'Florianópolis', false, MD5('9vYbW2')),
('Thiago Martins', 'thiago.martins@example.com', 'Frigorífico', 'Cuiabá', false, MD5('sK4nL8'));

INSERT INTO "tb_feedbacks" (employee_registry, content)
SELECT registry, 
	CASE 
    WHEN RANDOM() < 0.10 THEN 'Implementação de treinamentos especializados.'
    WHEN RANDOM() < 0.20 THEN 'Cultura de trabalho que não promove o equilíbrio vida-trabalho.'
		WHEN RANDOM() < 0.30 THEN 'Melhoria na comunicação interna.'
    WHEN RANDOM() < 0.40 THEN 'Aprimoramento nas condições de trabalho.'
		WHEN RANDOM() < 0.50 THEN 'Introdução de práticas de trabalho remoto e flexibilidade.'
    WHEN RANDOM() < 0.60 THEN 'Dificuldades na obtenção de feedback construtivo.'
		WHEN RANDOM() < 0.70 THEN 'Desafios ao equilibrar trabalho remoto e vida pessoal.'
    WHEN RANDOM() < 0.80 THEN 'Atualização e melhoria contínua da infraestrutura tecnológica.'
		WHEN RANDOM() < 0.90 THEN 'Problemas na coordenação entre equipes e setores.'
    WHEN RANDOM() < 1 THEN 'Ambiente de trabalho não sustentável e preocupações ecológicas.'
  END 
FROM "tb_employees" 
ORDER BY RANDOM()