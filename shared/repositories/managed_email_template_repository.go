package repositories

import "database/sql"

// ManagedEmailTemplateRepository handled DB operatios relative to the Managed Email Template repositories.
type ManagedEmailTemplateRepository struct {
	DB *sql.DB
}

// ManagedEmailTemplate contains partial information about a Managed Email Template entity.
type ManagedEmailTemplate struct {
	ID 			int64
	ProviderID  int64
	InternalID  string
}

// FindTemplateByInternalID returns a pointer of ManagedEmailTemplate for the given internal id.
func (repo *ManagedEmailTemplateRepository) FindTemplateByInternalID(internalID string) (*ManagedEmailTemplate, error) {
	var managedTemplate ManagedEmailTemplate
	row := repo.DB.QueryRow(`
SELECT
	t.id template_id,
	t.provider_id provider_id,
	t.internal_id internal_id
FROM 
     ht_managed_email_template t
WHERE
	internal_id = ?
	`,
	internalID)

	switch err := row.Scan(&managedTemplate.ID, &managedTemplate.ProviderID, &managedTemplate.InternalID); err {
	case sql.ErrNoRows:
		return nil, nil
	case nil:
		return &managedTemplate, nil
	default:
		return nil, err

	}
}