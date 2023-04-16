package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Masterminds/squirrel"
	log "github.com/sirupsen/logrus"

	"github.com/mlhan1993/KongInterview/internal/service"
)

type Kong struct {
	db *sql.DB
}

func NewKong(db *sql.DB) (*Kong, error) {
	s := &Kong{
		db: db,
	}
	return s, nil
}

func (s *Kong) GetServices(ctx context.Context, numPerPage, pageNumber uint, sortOrder, filter string) (uint, []service.Service, error) {
	var services []service.Service

	query := getServicesQuery(numPerPage, pageNumber, sortOrder, filter)

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return 0, nil, err
	}

	defer rows.Close()

	var total uint
	for rows.Next() {
		var myService service.Service
		//
		if err := rows.Scan(&myService.ID, &myService.Name, &myService.Description, &myService.NumVersions, &total); err != nil {
			return 0, nil, err
		}

		services = append(services, myService)
	}

	if err := rows.Err(); err != nil {
		return 0, nil, err
	}

	return total, services, nil
}

func (s *Kong) GetVersions(ctx context.Context, serviceId, numPerPage, pageNumber uint, sortOrder string) (uint, []service.Version, error) {
	// Implement GetVersions method
	var versions []service.Version

	query := getVersionsQuery(serviceId, numPerPage, pageNumber, sortOrder)
	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return 0, nil, err
	}

	defer rows.Close()

	var total uint
	for rows.Next() {
		var myVersion service.Version
		var dateCreated string

		if err := rows.Scan(&myVersion.ID, &myVersion.Tag, &myVersion.ServiceID, &dateCreated, &total); err != nil {
			return 0, nil, err
		}
		myVersion.DateCreated, err = time.Parse("2006-01-02 15:04:05", dateCreated)
		if err != nil {
			return 0, nil, err
		}
		versions = append(versions, myVersion)
	}
	return total, versions, nil
}

func getServicesQuery(numPerPage, pageNumber uint, sortOrder string, filter string) string {
	totalQuery := squirrel.Select("*, COUNT(*) OVER () AS total").From("services")
	if filter != "" {
		totalQuery = totalQuery.Where(fmt.Sprintf("name LIKE \"%%%s%%\" OR description LIKE \"%%%s%%\"",
			filter, filter))
	}

	sb := squirrel.Select("s.id", "s.name", "s.description", "COUNT(v.id) AS version_count", "s.total").
		FromSelect(totalQuery, "s").
		LeftJoin("versions v ON s.id = v.serviceID").
		GroupBy("s.id", "s.name", "s.description", "total")

	if sortOrder == "desc" {
		sb = sb.OrderBy("s.name DESC")
	} else {
		sb = sb.OrderBy("s.name ASC")
	}

	if numPerPage > 0 {
		sb = sb.Offset(uint64((pageNumber - 1) * numPerPage)).Limit(uint64(numPerPage))
	}

	s, _, _ := sb.ToSql()
	log.Debug(s)
	return s
}

func getVersionsQuery(serviceID, numPerPage, pageNumber uint, sortOrder string) string {

	totalQuery := squirrel.Select("Count(*) as total").From("versions").
		Where(fmt.Sprintf("serviceID = %d", serviceID))
	totalQueryStr, _, _ := totalQuery.ToSql()

	sb := squirrel.Select("v.id, v.tag, v.serviceID, v.dateCreated, total_rows.total").
		From("versions v").
		CrossJoin(fmt.Sprintf("(%s) AS total_rows", totalQueryStr)).
		Where(fmt.Sprintf("serviceID = %d", serviceID))

	if sortOrder == "desc" {
		sb = sb.OrderBy("dateCreated desc")
	} else {
		sb = sb.OrderBy("dateCreated asc")
	}
	if pageNumber > 0 {
		sb = sb.Limit(uint64(numPerPage)).Offset(uint64(numPerPage * (pageNumber - 1)))
	}
	s, _, _ := sb.ToSql()
	log.Debug(s)
	return s
}
