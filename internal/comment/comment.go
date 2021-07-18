package comment

import (
  "fmt"
  // "strconv"

  "github.com/jinzhu/gorm"
)

// Service - the struct for our comment service
type Service struct {
  DB *gorm.DB
}

// Comment - defines our comment structure
type Comment struct {
  gorm.Model
  Slug string
  Body string
  Author string
}

// CommentService - the interface for our comment service
type CommentService interface{
  GetComment(ID uint) (Comment, error)
  GetCommentsBySlug(slug string) ([]Comment, error)
  PostComment(commen Comment) (Comment, error)
  UpdateComment(ID uint, newComment Comment) (Comment, error)
  DeleteComment(ID uint) error
  GetAllComments() ([]Comment, error) 
}

// NewService - returns a comment service
func NewService(db *gorm.DB) *Service {
  return &Service{
   DB: db,
  }
}

// GetComment - retrieves comments by their ID from the database
func (s *Service) GetComment(ID uint) (Comment, error) {
  var comment Comment
  fmt.Printf("GetComment by ID: %d\n", ID)
  if result := s.DB.First(&comment, ID); result.Error != nil {
    fmt.Printf("Returning empty comment.\n")
    return Comment{}, result.Error
  }
  return comment, nil
} 

// GetCommentsBySlug - retrieves all comments by slug (path - /article/name/)
func (s *Service)  GetCommentsBySlug(slug string) ([]Comment, error) {
  var comments []Comment
  if result := s.DB.Find(&comments).Where("slug = ?", slug); result.Error != nil {
    return []Comment{}, result.Error
  }
  return comments, nil
}

// PostComment - adds a new comment to the database
func (s *Service) PostComment(comment Comment) (Comment, error) {
  if result := s.DB.Save(&comment); result.Error != nil {
    return Comment{}, result.Error
  }
  return comment, nil
}

// UpdateComment - updates a comment by ID with new comment info
func (s *Service) UpdateComment(ID uint, newComment Comment) (Comment, error) {
  fmt.Printf("UpdateComment by ID: %d\n", ID)
  comment, err := s.GetComment(ID)
  if err != nil {
    fmt.Printf("Comment by ID: %d - not found.\n", ID)
    return Comment{}, err
  }

  if result := s.DB.Model(&comment).Updates(newComment); result.Error != nil {
    return Comment{}, result.Error
  }

  return comment, nil
}

// DeleteComment - deletes a comment from the database by ID
func (s *Service)  DeleteComment(ID uint) error {
  fmt.Printf("DeleteComment by ID: %d\n", ID)

  _, err := s.GetComment(ID)

  if err != nil {
    fmt.Println("Comment for this ID does not exist.")
    return err
  }

  if result := s.DB.Delete(&Comment{}, ID); result.Error != nil {
    return result.Error
  }
  return nil
}

// GetAllComments - retrieves all comments from the database
func (s *Service)  GetAllComments() ([]Comment, error) {
  var comments []Comment
  if result := s.DB.Find(&comments); result.Error != nil {
    return comments, result.Error
  }
  return comments, nil
}
