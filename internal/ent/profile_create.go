// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"go_plate/internal/ent/account"
	"go_plate/internal/ent/profile"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProfileCreate is the builder for creating a Profile entity.
type ProfileCreate struct {
	config
	mutation *ProfileMutation
	hooks    []Hook
}

// SetUsername sets the "username" field.
func (pc *ProfileCreate) SetUsername(s string) *ProfileCreate {
	pc.mutation.SetUsername(s)
	return pc
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (pc *ProfileCreate) SetNillableUsername(s *string) *ProfileCreate {
	if s != nil {
		pc.SetUsername(*s)
	}
	return pc
}

// SetDisplayName sets the "display_name" field.
func (pc *ProfileCreate) SetDisplayName(s string) *ProfileCreate {
	pc.mutation.SetDisplayName(s)
	return pc
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (pc *ProfileCreate) SetNillableDisplayName(s *string) *ProfileCreate {
	if s != nil {
		pc.SetDisplayName(*s)
	}
	return pc
}

// SetBio sets the "bio" field.
func (pc *ProfileCreate) SetBio(s string) *ProfileCreate {
	pc.mutation.SetBio(s)
	return pc
}

// SetNillableBio sets the "bio" field if the given value is not nil.
func (pc *ProfileCreate) SetNillableBio(s *string) *ProfileCreate {
	if s != nil {
		pc.SetBio(*s)
	}
	return pc
}

// SetLangTag sets the "lang_tag" field.
func (pc *ProfileCreate) SetLangTag(s string) *ProfileCreate {
	pc.mutation.SetLangTag(s)
	return pc
}

// SetNillableLangTag sets the "lang_tag" field if the given value is not nil.
func (pc *ProfileCreate) SetNillableLangTag(s *string) *ProfileCreate {
	if s != nil {
		pc.SetLangTag(*s)
	}
	return pc
}

// SetLocation sets the "location" field.
func (pc *ProfileCreate) SetLocation(s string) *ProfileCreate {
	pc.mutation.SetLocation(s)
	return pc
}

// SetNillableLocation sets the "location" field if the given value is not nil.
func (pc *ProfileCreate) SetNillableLocation(s *string) *ProfileCreate {
	if s != nil {
		pc.SetLocation(*s)
	}
	return pc
}

// SetTimezone sets the "timezone" field.
func (pc *ProfileCreate) SetTimezone(s string) *ProfileCreate {
	pc.mutation.SetTimezone(s)
	return pc
}

// SetNillableTimezone sets the "timezone" field if the given value is not nil.
func (pc *ProfileCreate) SetNillableTimezone(s *string) *ProfileCreate {
	if s != nil {
		pc.SetTimezone(*s)
	}
	return pc
}

// SetUpdateTime sets the "update_time" field.
func (pc *ProfileCreate) SetUpdateTime(t time.Time) *ProfileCreate {
	pc.mutation.SetUpdateTime(t)
	return pc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (pc *ProfileCreate) SetNillableUpdateTime(t *time.Time) *ProfileCreate {
	if t != nil {
		pc.SetUpdateTime(*t)
	}
	return pc
}

// SetCreateTime sets the "create_time" field.
func (pc *ProfileCreate) SetCreateTime(t time.Time) *ProfileCreate {
	pc.mutation.SetCreateTime(t)
	return pc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (pc *ProfileCreate) SetNillableCreateTime(t *time.Time) *ProfileCreate {
	if t != nil {
		pc.SetCreateTime(*t)
	}
	return pc
}

// SetID sets the "id" field.
func (pc *ProfileCreate) SetID(s string) *ProfileCreate {
	pc.mutation.SetID(s)
	return pc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (pc *ProfileCreate) SetNillableID(s *string) *ProfileCreate {
	if s != nil {
		pc.SetID(*s)
	}
	return pc
}

// SetAccountID sets the "account" edge to the Account entity by ID.
func (pc *ProfileCreate) SetAccountID(id string) *ProfileCreate {
	pc.mutation.SetAccountID(id)
	return pc
}

// SetAccount sets the "account" edge to the Account entity.
func (pc *ProfileCreate) SetAccount(a *Account) *ProfileCreate {
	return pc.SetAccountID(a.ID)
}

// Mutation returns the ProfileMutation object of the builder.
func (pc *ProfileCreate) Mutation() *ProfileMutation {
	return pc.mutation
}

// Save creates the Profile in the database.
func (pc *ProfileCreate) Save(ctx context.Context) (*Profile, error) {
	pc.defaults()
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *ProfileCreate) SaveX(ctx context.Context) *Profile {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *ProfileCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *ProfileCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *ProfileCreate) defaults() {
	if _, ok := pc.mutation.LangTag(); !ok {
		v := profile.DefaultLangTag
		pc.mutation.SetLangTag(v)
	}
	if _, ok := pc.mutation.UpdateTime(); !ok {
		v := profile.DefaultUpdateTime()
		pc.mutation.SetUpdateTime(v)
	}
	if _, ok := pc.mutation.CreateTime(); !ok {
		v := profile.DefaultCreateTime()
		pc.mutation.SetCreateTime(v)
	}
	if _, ok := pc.mutation.ID(); !ok {
		v := profile.DefaultID()
		pc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *ProfileCreate) check() error {
	if _, ok := pc.mutation.LangTag(); !ok {
		return &ValidationError{Name: "lang_tag", err: errors.New(`ent: missing required field "Profile.lang_tag"`)}
	}
	if _, ok := pc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "Profile.update_time"`)}
	}
	if _, ok := pc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "Profile.create_time"`)}
	}
	if v, ok := pc.mutation.ID(); ok {
		if err := profile.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "Profile.id": %w`, err)}
		}
	}
	if _, ok := pc.mutation.AccountID(); !ok {
		return &ValidationError{Name: "account", err: errors.New(`ent: missing required edge "Profile.account"`)}
	}
	return nil
}

func (pc *ProfileCreate) sqlSave(ctx context.Context) (*Profile, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Profile.ID type: %T", _spec.ID.Value)
		}
	}
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *ProfileCreate) createSpec() (*Profile, *sqlgraph.CreateSpec) {
	var (
		_node = &Profile{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(profile.Table, sqlgraph.NewFieldSpec(profile.FieldID, field.TypeString))
	)
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := pc.mutation.Username(); ok {
		_spec.SetField(profile.FieldUsername, field.TypeString, value)
		_node.Username = value
	}
	if value, ok := pc.mutation.DisplayName(); ok {
		_spec.SetField(profile.FieldDisplayName, field.TypeString, value)
		_node.DisplayName = value
	}
	if value, ok := pc.mutation.Bio(); ok {
		_spec.SetField(profile.FieldBio, field.TypeString, value)
		_node.Bio = value
	}
	if value, ok := pc.mutation.LangTag(); ok {
		_spec.SetField(profile.FieldLangTag, field.TypeString, value)
		_node.LangTag = value
	}
	if value, ok := pc.mutation.Location(); ok {
		_spec.SetField(profile.FieldLocation, field.TypeString, value)
		_node.Location = value
	}
	if value, ok := pc.mutation.Timezone(); ok {
		_spec.SetField(profile.FieldTimezone, field.TypeString, value)
		_node.Timezone = value
	}
	if value, ok := pc.mutation.UpdateTime(); ok {
		_spec.SetField(profile.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	if value, ok := pc.mutation.CreateTime(); ok {
		_spec.SetField(profile.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if nodes := pc.mutation.AccountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   profile.AccountTable,
			Columns: []string{profile.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.account_profile = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ProfileCreateBulk is the builder for creating many Profile entities in bulk.
type ProfileCreateBulk struct {
	config
	err      error
	builders []*ProfileCreate
}

// Save creates the Profile entities in the database.
func (pcb *ProfileCreateBulk) Save(ctx context.Context) ([]*Profile, error) {
	if pcb.err != nil {
		return nil, pcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Profile, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProfileMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *ProfileCreateBulk) SaveX(ctx context.Context) []*Profile {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *ProfileCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *ProfileCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
