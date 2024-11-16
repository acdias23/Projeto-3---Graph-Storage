package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/jaswdr/faker"
	"github.com/joho/godotenv"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type Aluno struct {
	Name string
	Date time.Time
}

var cursos = []string{
	"Administracao",
	"Ciencia da Computacao",
	"Engenharia Civil",
	"Engenharia Eletrica",
}

var cursos_id = [4]string{}
var departamentos_id = [5]string{}
var disciplinas_adm_id = [40]string{}
var disciplinas_cc_id = [40]string{}
var disciplinas_engc_id = [40]string{}
var disciplinas_enge_id = [40]string{}
var alunos_id = [20]string{}
var professores_id = [15]string{}

func pickRandom(arr []string, n int) string {
	idx := rand.Intn(n)
	return arr[idx]
}

func randYear() int {
	min := time.Date(2010, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2020, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0).Year()
}

func main() {
	ctx := context.Background()
	// URI examples: "neo4j://localhost", "neo4j+s://xxx.databases.neo4j.io"
	_ = godotenv.Load()
	dbUri := os.Getenv("NEO4J_URI")
	dbUser := os.Getenv("NEO4J_USERNAME")
	dbPassword := os.Getenv("NEO4J_PASSWORD")

	driver, err := neo4j.NewDriverWithContext(
		dbUri,
		neo4j.BasicAuth(dbUser, dbPassword, ""))
	defer driver.Close(ctx)

	err = driver.VerifyConnectivity(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connection established.")

	m := make(map[string][]string)

	m["Administracao"] = []string{
		"Matematica aplicada a administracao",
		"Fundamentos da administracao",
		"Estudos em macroeconomia",
		"Sociologia I",
		"Sociologia II",
		"Linguagens e generos textuais",
		"Etica nas organizacoes",
		"Ensino social cristao",
		"Calculo basico I",
		"Calculo basico II",
		"Modelos organizacionais",
		"Contabilidade financeira",
		"Estudos em microeconomia",
		"Filosofia I",
		"Filosofia II",
		"Linguagem em comunicacao organizacional",
		"Calculo diferencial e integral I",
		"Calculo diferencial e integral II",
		"Calculo diferencial e integral III",
		"Calculo diferencial e integral IV",
		"Gramatica",
		"Estudos em microeconomia",
		"Sustentabilidade",
		"Probabilidade estatistica",
		"Estudos do MEI",
		"Gestao de negocios",
		"Gestao de pessoas",
		"Gesta Financeira",
		"Matematica Financeira",
		"Gestao de recusos humanos",
		"PowerBI",
		"Excel Basico",
		"Excel intermediario",
		"Excel avancado",
		"Redacao",
		"Comunicacao para empresas",
		"Estrategia de marketing",
		"Direito Tributario",
		"Logistica",
		"Implementacao de negocios",
	}

	m["Ciencia da Computacao"] = []string{
		"Filosofia",
		"Linguagem em comunicacao organizacional",
		"Calculo diferencial e integral I",
		"Fundamentos de algoritmos",
		"Desenvolvimento web",
		"Comunicacao e expressao",
		"Calculo diferencial e integral II",
		"Calculo diferencial e integral III",
		"Calculo diferencial e integral IV",
		"Geometria analitica",
		"Teoria dos grafos",
		"Desenvolvimento de algoritmos",
		"Fisica",
		"Sociologia",
		"Praticas de inovacao",
		"Eletronica geral",
		"Quimica",
		"Circuitos eletricos",
		"Ecologia e sustentabilidade",
		"Sistemas embarcados",
		"Redes de computadores",
		"Calculo numerico",
		"Termodinamica",
		"Cinematica e dinamica",
		"Algebra Linear",
		"Automatos",
		"Redes moveis",
		"Algoritmos",
		"Sistema Operacional",
		"Computacao grafica",
		"Compiladores",
		"Empreendedorismo",
		"Inovacao",
		"Teste de software",
		"IA",
		"Orientacao a objetos",
		"Desenvolvimento de jogos",
		"Robotica",
		"TCC I",
		"TCC II",
	}

	m["Engenharia Civil"] = []string{
		"Ensino social cristao",
		"Calculo basico",
		"Modelos organizacionais",
		"Contabilidade financeira",
		"Estudos em microeconomia",
		"Filosofia",
		"Linguagem em comunicacao organizacional",
		"Calculo diferencial e integral I",
		"Fundamentos de algoritmos",
		"Desenvolvimento web",
		"Comunicacao e expressao",
		"Calculo diferencial e integral II",
		"Calculo diferencial e integral III",
		"Calculo diferencial e integral IV",
		"Geometria analitica",
		"Teoria dos grafos",
		"Desenvolvimento de algoritmos",
		"Fisica",
		"Desenho tecnico",
		"Praticas de inovacao",
		"Eletronica geral",
		"Quimica",
		"Arquitetura e representacao grafica",
		"Mecanica geral",
		"Topografia",
		"Eletricidade geral",
		"Instalacoes eletricas",
		"Economia",
		"Geotecnia I",
		"Geotecnia II",
		"Transportes I",
		"Transportes II",
		"Optativa I",
		"Optativa II",
		"Custos",
		"Tecnologia das construcoes",
		"Metodos estatisticos",
		"Termodinamica",
		"Gestao de obras",
		"Planejamento de obras",
	}

	m["Engenharia Eletrica"] = []string{
		"Ensino social cristao",
		"Calculo basico",
		"Modelos organizacionais",
		"Contabilidade financeira",
		"Estudos em microeconomia",
		"Filosofia",
		"Linguagem em comunicacao organizacional",
		"Calculo diferencial e integral I",
		"Fundamentos de algoritmos",
		"Desenvolvimento web",
		"Comunicacao e expressao",
		"Calculo diferencial e integral II",
		"Calculo diferencial e integral III",
		"Calculo diferencial e integral IV",
		"Geometria analitica",
		"Teoria dos grafos",
		"Desenvolvimento de algoritmos",
		"Fisica",
		"Desenho tecnico",
		"Praticas de inovacao",
		"Eletronica geral",
		"Quimica",
		"sinais e Sistemas",
		"Praticas de inovacao I",
		"Praticas de inovacao II",
		"Praticas de inovacao III",
		"Circuitos integrados",
		"Mecanica Geral",
		"Gestao organizacional",
		"TCC 1",
		"TCC 2",
		"Redes",
		"Mecanica dos solidos",
		"Mecanica dos fluidos",
		"Sistemas eletricos",
		"Seguranca do trabalho",
		"Economia",
		"Conversao de energia I",
		"Conversao de energia II",
		"Conversao de energia III",
	}

	departamentos := [5]string{
		"Matematica",
		"Ciencia da Computacao",
		"Quimica",
		"Fisica",
		"Engenharia",
	}

	fakeName := faker.New().Person()
	fakeDate := faker.New().Time()

	var pk string
	for i := 0; i < 20; i++ {
		pk, _ = insertAluno(driver, Aluno{fakeName.Name(), fakeDate.Time(time.Now())})
		alunos_id[i] = pk
	}

	for i := 0; i < 15; i++ {
		pk, _ := insertProfessor(driver, fakeName.Name())
		professores_id[i] = pk
	}

	for i, curso := range cursos {
		pk, _ := insertCurso(driver, curso)
		cursos_id[i] = pk
	}

	for i, departamento := range departamentos {
		pk, _ := insertDepartamento(driver, departamento)
		departamentos_id[i] = pk
	}

	for i, disc := range m["Administracao"] {
		pk, _ := insertDisciplina(driver, disc)
		disciplinas_adm_id[i] = pk
	}

	for i, disc := range m["Ciencia da Computacao"] {
		pk, _ := insertDisciplina(driver, disc)
		disciplinas_cc_id[i] = pk
	}

	for i, disc := range m["Engenharia Civil"] {
		pk, _ := insertDisciplina(driver, disc)
		disciplinas_engc_id[i] = pk
	}

	for i, disc := range m["Engenharia Eletrica"] {
		pk, _ := insertDisciplina(driver, disc)
		disciplinas_enge_id[i] = pk
	}

	atualizarFKAluno(driver)
	atualizarFKCurso(driver)
	atualizarFKDepartamento(driver)
	atualizarFKDisciplina(driver, disciplinas_adm_id[:], 0)
	atualizarFKDisciplina(driver, disciplinas_cc_id[:], 1)
	atualizarFKDisciplina(driver, disciplinas_engc_id[:], 2)
	atualizarFKDisciplina(driver, disciplinas_enge_id[:], 3)
	atualizarFKProfessor(driver)
	atualizarDisciplinasMinistradas(ctx, driver, disciplinas_adm_id[:])
	atualizarDisciplinasMinistradas(ctx, driver, disciplinas_cc_id[:])
	atualizarDisciplinasMinistradas(ctx, driver, disciplinas_engc_id[:])
	atualizarDisciplinasMinistradas(ctx, driver, disciplinas_enge_id[:])
	criarGruposTCC(ctx, driver)
	criarHistoricoEscolar(ctx, driver)
	criarMatrizCurricular(ctx, driver)
}

func atualizarDisciplinasMinistradas(ctx context.Context, driver neo4j.DriverWithContext, disciplinas []string) {
	// Inicialize uma sessão de escrita com contexto
	session := driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	for i, disc := range disciplinas {
		professorID := pickRandom(professores_id[:], len(professores_id))
		semestre := int(i/5) + 1
		ano := randYear()

		// Execute a transação de inserção com contexto
		_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
			query := `
				CREATE (d:DisciplinaMinistrada {
					professor_id: $professorID,
					disciplina_id: $disciplinaID,
					semestre: $semestre,
					ano: $ano
				})
			`
			params := map[string]interface{}{
				"professorID":  professorID,
				"disciplinaID": disc,
				"semestre":     semestre,
				"ano":          ano,
			}

			_, err := tx.Run(ctx, query, params)
			return nil, err
		})

		if err != nil {
			log.Fatal(err)
		}
	}
}

func criarGruposTCC(ctx context.Context, driver neo4j.DriverWithContext) {
	session := driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	for i := 0; i < 6; i++ {
		aluno1 := pickRandom(alunos_id[:], len(alunos_id))
		aluno2 := pickRandom(alunos_id[:], len(alunos_id))
		aluno3 := pickRandom(alunos_id[:], len(alunos_id))
		orientador := pickRandom(professores_id[:], len(professores_id))

		// Verifica se os IDs retornados são válidos e exibe logs

		if aluno1 == "" || aluno2 == "" || aluno3 == "" || orientador == "" {
			log.Fatal("Erro: ID inválido retornado por pickRandom.")
		}

		_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
			query := `
				CREATE (g:GrupoTCC {
					aluno1_id: $aluno1,
					aluno2_id: $aluno2,
					aluno3_id: $aluno3,
					orientador: $orientador
				})
			`
			params := map[string]interface{}{
				"aluno1":     aluno1,
				"aluno2":     aluno2,
				"aluno3":     aluno3,
				"orientador": orientador,
			}
			//println(aluno1, aluno2, aluno3)
			_, err := tx.Run(ctx, query, params)
			return nil, err
		})

		if err != nil {
			log.Fatal("Erro ao inserir grupo de TCC: ", err)
		}
	}
}

func criarMatrizCurricular(ctx context.Context, driver neo4j.DriverWithContext) {
	for _, curso := range cursos_id {
		// Map disciplinas based on the course
		var d [40]string
		switch curso {
		case cursos_id[0]:
			d = disciplinas_adm_id
		case cursos_id[1]:
			d = disciplinas_cc_id
		case cursos_id[2]:
			d = disciplinas_engc_id
		case cursos_id[3]:
			d = disciplinas_enge_id
		}

		session := driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
		// Start a transaction to insert the relationships
		_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
			// Loop over the disciplinas to create the relationships
			for i, disc := range d {
				// Create the relationship between Curso and Disciplina
				cypherQuery := `
					MATCH (c:Curso {curso_id: $cursoId}), (d:Disciplina {disciplina_id: $discId})
					CREATE (c)-[:OFERECE {semestre: $semestre}]->(d)
				`
				semester := int(i/5) + 1 // Mapping semester based on index
				_, err := tx.Run(ctx, cypherQuery, map[string]interface{}{
					"cursoId":  curso,
					"discId":   disc,
					"semestre": semester,
				})
				if err != nil {
					return nil, err
				}
			}
			return nil, nil
		})

		if err != nil {
			log.Fatalf("Error inserting data into matriz_curricular for curso %v: %v", curso, err)
		}
	}

}

func criarHistoricoEscolar(ctx context.Context, driver neo4j.DriverWithContext) {
	// Open a session
	session := driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	// Process each student
	for _, aluno := range alunos_id {
		var dataEntrada string
		var curso string

		record, err := session.Run(ctx,
			`MATCH (a:Aluno {aluno_id: $alunoId})-[r:ESTUDA_EM]->(c:Curso)
				 RETURN a.data_ingresso AS dataEntrada, c.curso_id AS cursoId`,
			map[string]interface{}{"alunoId": aluno},
		)
		if err != nil {
			return
		}

		// Check if a record was returned
		if record.Next(ctx) {
			rec := record.Record()
			data, _ := rec.Get("dataEntrada")
			cursoid, _ := rec.Get("cursoId")
			dataEntrada = data.(string)
			curso = cursoid.(string)
		}

		// Extract the year from the entry date
		layout := "2006-01-02T15:04:05-07:00" // Layout for the given format
		data, _ := time.Parse(layout, dataEntrada)
		ano := int64(data.Year())

		// Determine the disciplines for the student's course
		var disciplinas [40]string
		switch curso {
		case cursos_id[0]:
			disciplinas = disciplinas_adm_id
		case cursos_id[1]:
			disciplinas = disciplinas_cc_id
		case cursos_id[2]:
			disciplinas = disciplinas_engc_id
		case cursos_id[3]:
			disciplinas = disciplinas_enge_id
		}

		// Insert data for each discipline
		for i, disc := range disciplinas {
			semestre := int(i/5) + 1
			anoFinal := ano + int64(i/10)
			notaFinal := rand.Float32() * 10

			_, err := session.Run(ctx,
				`MATCH (a:Aluno {aluno_id: $alunoId})
                     MATCH (d:Disciplina {disciplina_id: $disciplinaId})
                     CREATE (a)-[:CURSOU {
                         semestre: $semestre,
                         ano: $ano,
                         nota_final: $notaFinal
                     }]->(d)`,
				map[string]interface{}{
					"alunoId":      aluno,
					"disciplinaId": disc,
					"semestre":     semestre,
					"ano":          anoFinal,
					"notaFinal":    notaFinal,
				})

			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func insertAluno(driver neo4j.DriverWithContext, aluno Aluno) (string, error) {
	// Criar o contexto com timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Criar uma sessão com o contexto
	session := driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	// Definir o Cypher para criar um nó de "Aluno"
	cypherQuery := `
		CREATE (a:Aluno {aluno_id: $id, nome: $nome, data_ingresso: $data_ingresso})
		RETURN a.aluno_id AS alunoId
	`

	// Criar parâmetros para a consulta Cypher
	params := map[string]interface{}{
		"nome":          aluno.Name,
		"data_ingresso": aluno.Date.Format(time.RFC3339),
		"id":            uuid.New().String(),
	}

	// Executar a consulta com o contexto
	result, err := session.Run(ctx, cypherQuery, params)
	if err != nil {
		log.Fatal("Erro ao rodar a consulta:", err)
	}

	// Obter o ID do nó criado
	if result.Next(ctx) {
		id, ok := result.Record().Get("alunoId")
		if !ok {
			return "", fmt.Errorf("falha ao converter ID do aluno")
		}
		return id.(string), nil
	}

	log.Fatal("Erro ao obter o ID do aluno.")
	return "", result.Err()
}

func atualizarFKAluno(driver neo4j.DriverWithContext) error {
	// Criar um contexto com timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Abrir uma sessão de escrita no Neo4j
	session := driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	// Iterar sobre cada aluno
	for _, alunoID := range alunos_id {
		// Selecionar um curso aleatório
		newCursoID := pickRandom(cursos_id[:], len(cursos_id))

		// Executar uma transação para atualizar o relacionamento
		_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
			// Remover qualquer relacionamento `ESTUDA_EM` existente para o aluno
			_, err := tx.Run(ctx, `
                MATCH (a:Aluno)-[r:ESTUDA_EM]->(:Curso)
                WHERE a.aluno_id = $alunoID
                DELETE r
            `, map[string]any{"alunoID": alunoID})
			if err != nil {
				return nil, err
			}

			// Criar um novo relacionamento `ESTUDA_EM` entre o aluno e o novo curso
			_, err = tx.Run(ctx, `
                MATCH (a:Aluno), (c:Curso)
                WHERE a.aluno_id = $alunoID AND c.curso_id = $cursoID
                CREATE (a)-[:ESTUDA_EM]->(c)
            `, map[string]any{"alunoID": alunoID, "cursoID": newCursoID})
			return nil, err
		})

		if err != nil {
			log.Printf("Erro ao atualizar o curso para o aluno %s: %v", alunoID, err)
			return err
		}
	}

	return nil
}

func insertProfessor(driver neo4j.DriverWithContext, nome string) (string, error) {
	// Criar o contexto com timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Criar uma sessão com o contexto
	session := driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	// Definir o Cypher para criar um nó de "Aluno"
	cypherQuery := `
		CREATE (p:Professor {professor_id: $id, nome: $nome})
		RETURN p.professor_id AS professorId
	`

	// Criar parâmetros para a consulta Cypher
	params := map[string]interface{}{
		"nome": nome,
		"id":   uuid.New().String(),
	}

	// Executar a consulta com o contexto
	result, err := session.Run(ctx, cypherQuery, params)
	if err != nil {
		log.Fatal("Erro ao rodar a consulta:", err)
	}

	// Obter o ID do nó criado
	if result.Next(ctx) {
		id, ok := result.Record().Get("professorId")
		if !ok {
			return "", fmt.Errorf("falha ao converter ID do professor")
		}
		return id.(string), nil
	}

	log.Fatal("Erro ao obter o ID do professor.")
	return "", result.Err()
}

func atualizarFKProfessor(driver neo4j.DriverWithContext) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	session := driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	for _, professorID := range professores_id {
		newDepartamentoID := pickRandom(departamentos_id[:], len(departamentos_id))

		_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
			// Remover o relacionamento `PERTENCE_A` existente do professor com qualquer departamento
			_, err := tx.Run(ctx, `
                MATCH (p:Professor)-[r:PERTENCE_A]->(:Departamento)
                WHERE p.professor_id = $professorID
                DELETE r
            `, map[string]any{"professorID": professorID})
			if err != nil {
				return nil, err
			}

			// Criar um novo relacionamento `PERTENCE_A` com o novo departamento
			_, err = tx.Run(ctx, `
                MATCH (p:Professor), (d:Departamento)
                WHERE p.professor_id = $professorID AND d.departamento_id = $departamentoID
                CREATE (p)-[:PERTENCE_A]->(d)
            `, map[string]any{"professorID": professorID, "departamentoID": newDepartamentoID})
			return nil, err
		})

		if err != nil {
			log.Printf("Erro ao atualizar o departamento para o professor %s: %v", professorID, err)
			return err
		}
	}

	return nil
}

func insertCurso(driver neo4j.DriverWithContext, nome string) (string, error) {
	// Criar um contexto com timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Abrir uma nova sessão
	session := driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	// Executar a transação para criar o nó do curso
	cursoID, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		// Executar a consulta para criar um nó de curso
		result, err := tx.Run(ctx, "CREATE (c:Curso {curso_id: $id, nome: $nome}) RETURN c.curso_id AS cursoID",
			map[string]any{"id": uuid.New().String(), "nome": nome})
		if err != nil {
			return nil, err
		}

		// Obter o ID do nó criado
		if result.Next(ctx) {
			id, ok := result.Record().Get("cursoID")
			if !ok {
				return nil, fmt.Errorf("falha ao converter ID do curso")
			}
			return id, nil
		}

		return nil, result.Err()
	})

	if err != nil {
		return "", fmt.Errorf("erro ao criar curso: %w", err)
	}

	// Retornar o ID do nó criado
	return cursoID.(string), nil
}

func atualizarFKCurso(driver neo4j.DriverWithContext) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	session := driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	for _, cursoID := range cursos_id {
		// Seleciona aleatoriamente um ID de departamento
		newDepartamentoID := pickRandom(departamentos_id[:], len(departamentos_id))

		_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
			// Remover o relacionamento `PERTENCE_A` existente do curso com qualquer departamento
			_, err := tx.Run(ctx, `
                MATCH (c:Curso)-[r:PERTENCE_A]->(:Departamento)
                WHERE c.curso_id = $cursoID
                DELETE r
            `, map[string]any{"cursoID": cursoID})
			if err != nil {
				return nil, err
			}

			// Criar um novo relacionamento `PERTENCE_A` entre curso e departamento
			_, err = tx.Run(ctx, `
                MATCH (c:Curso), (d:Departamento)
                WHERE c.curso_id = $cursoID AND d.departamento_id = $departamentoID
                CREATE (c)-[:PERTENCE_A]->(d)
            `, map[string]any{"cursoID": cursoID, "departamentoID": newDepartamentoID})
			return nil, err
		})

		if err != nil {
			log.Printf("Erro ao atualizar o departamento para o curso %s: %v", cursoID, err)
			return err
		}
	}

	return nil
}

func insertDepartamento(driver neo4j.DriverWithContext, nome string) (string, error) {
	// Criar um contexto com timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Criar uma sessão para interagir com o banco de dados
	session := driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	cypherQuery := `
	CREATE (d:Departamento {departamento_id: $id, nome: $nome})
	RETURN d.departamento_id as departamentoId
`

	// Criar parâmetros para a consulta Cypher
	params := map[string]interface{}{
		"nome": nome,
		"id":   uuid.New().String(),
	}

	// Executar a consulta com o contexto
	result, err := session.Run(ctx, cypherQuery, params)
	if err != nil {
		log.Fatal("Erro ao rodar a consulta:", err)
	}

	// Obter o ID do nó criado
	if result.Next(ctx) {
		id, ok := result.Record().Get("departamentoId")
		if !ok {
			return "", fmt.Errorf("falha ao converter ID do departamento")
		}
		return id.(string), nil
	}

	log.Fatal("Erro ao obter o ID do departamento.")
	return "", result.Err()
}

func atualizarFKDepartamento(driver neo4j.DriverWithContext) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	session := driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	for _, departamentoID := range departamentos_id {
		// Seleciona aleatoriamente um ID de professor para ser o chefe do departamento
		newChefeID := pickRandom(professores_id[:], len(professores_id))

		_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
			// Remove o relacionamento `CHEFIA` existente do departamento com qualquer professor
			_, err := tx.Run(ctx, `
                MATCH (d:Departamento)-[r:CHEFIA]->(:Professor)
                WHERE d.departamento_id = $departamentoID
                DELETE r
            `, map[string]any{"departamentoID": departamentoID})
			if err != nil {
				return nil, err
			}

			// Cria um novo relacionamento `CHEFIA` entre o departamento e o novo professor chefe
			_, err = tx.Run(ctx, `
                MATCH (d:Departamento), (p:Professor)
                WHERE d.departamento_id = $departamentoID AND p.professor_id = $chefeID
                CREATE (d)-[:CHEFIA]->(p)
            `, map[string]any{"departamentoID": departamentoID, "chefeID": newChefeID})
			return nil, err
		})

		if err != nil {
			log.Printf("Erro ao atualizar o chefe do departamento %s: %v", departamentoID, err)
			return err
		}
	}

	return nil
}

func insertDisciplina(driver neo4j.DriverWithContext, nome string) (string, error) {
	// Criar um contexto com timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Criar uma sessão para interagir com o banco de dados
	session := driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	cypherQuery := `
	CREATE (d:Disciplina {disciplina_id: $id, nome: $nome})
	RETURN d.disciplina_id AS disciplinaId
`

	// Criar parâmetros para a consulta Cypher
	params := map[string]interface{}{
		"nome": nome,
		"id":   uuid.New().String(),
	}

	// Executar a consulta com o contexto
	result, err := session.Run(ctx, cypherQuery, params)
	if err != nil {
		log.Fatal("Erro ao rodar a consulta:", err)
	}

	// Obter o ID do nó criado
	if result.Next(ctx) {
		id, ok := result.Record().Get("disciplinaId")
		if !ok {
			return "", fmt.Errorf("falha ao converter ID do disciplina")
		}
		return id.(string), nil
	}

	log.Fatal("Erro ao obter o ID do disciplina.")
	return "", result.Err()
}

func atualizarFKDisciplina(driver neo4j.DriverWithContext, disciplinasID []string, idx int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	session := driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	cursoID := cursos_id[idx] // Seleciona o curso com base no índice fornecido

	for _, disciplinaID := range disciplinasID {
		professorID := pickRandom(professores_id[:], len(professores_id)) // Seleciona um ID de professor aleatoriamente

		_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
			// Remove relacionamentos `PERTENCE_A` e `MINISTRADA_POR` existentes para a disciplina
			_, err := tx.Run(ctx, `
                MATCH (d:Disciplina)-[r:PERTENCE_A|MINISTRADA_POR]->()
                WHERE d.disciplina_id = $disciplinaID
                DELETE r
            `, map[string]any{"disciplinaID": disciplinaID})
			if err != nil {
				return nil, err
			}

			// Cria um relacionamento `PERTENCE_A` entre a disciplina e o curso
			_, err = tx.Run(ctx, `
                MATCH (d:Disciplina), (c:Curso)
                WHERE d.disciplina_id = $disciplinaID AND c.curso_id = $cursoID
                CREATE (d)-[:PERTENCE_A]->(c)
            `, map[string]any{"disciplinaID": disciplinaID, "cursoID": cursoID})
			if err != nil {
				return nil, err
			}

			// Cria um relacionamento `MINISTRADA_POR` entre a disciplina e o professor
			_, err = tx.Run(ctx, `
                MATCH (d:Disciplina), (p:Professor)
                WHERE d.disciplina_id = $disciplinaID AND p.professor_id = $professorID
                CREATE (d)-[:MINISTRADA_POR]->(p)
            `, map[string]any{"disciplinaID": disciplinaID, "professorID": professorID})
			return nil, err
		})

		if err != nil {
			log.Printf("Erro ao atualizar os relacionamentos para a disciplina %s: %v", disciplinaID, err)
			return err
		}
	}

	return nil
}
