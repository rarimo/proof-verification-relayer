package data

type MasterQ interface {
	New() MasterQ

	Proofs() ProofsQ

	Transaction(fn func(db MasterQ) error) error
}
